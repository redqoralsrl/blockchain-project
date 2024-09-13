package swap

import (
	"blockscan-go/internal/config"
	"blockscan-go/internal/core/common/abi/erc20"
	"blockscan-go/internal/core/common/abi/factory"
	"blockscan-go/internal/core/common/abi/pair"
	"blockscan-go/internal/core/common/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

type Callers struct {
	conf           *config.Config
	factoryAbi     *abi.ABI
	routerAbi      *abi.ABI
	pairAbi        *abi.ABI
	factoryAddress *NetworkOfCA
	routerAddress  *NetworkOfCA
	logger         *zap.Logger
}

type NetworkOfCA struct {
	TestNetworkAddress string
	MainNetworkAddress string
}

func NewCallers(conf *config.Config, fa *abi.ABI, ra *abi.ABI, pa *abi.ABI, fi *NetworkOfCA, ri *NetworkOfCA, logger *zap.Logger) *Callers {
	return &Callers{
		conf:           conf,
		factoryAbi:     fa,
		routerAbi:      ra,
		pairAbi:        pa,
		factoryAddress: fi,
		routerAddress:  ri,
		logger:         logger,
	}
}

type Swap struct {
	Index             int            `json:"index"`
	PairAddress       common.Address `json:"pair_address"`
	Token0Address     common.Address `json:"token_0_address"`
	Token0Symbol      string         `json:"token_0_symbol,omitempty"`
	Token0Name        string         `json:"token_0_name,omitempty"`
	Token0Decimals    uint8          `json:"token_0_decimals,omitempty"`
	Token0GetReserves string         `json:"token_0_get_reserves"`
	Token1Address     common.Address `json:"token_1_address"`
	Token1Symbol      string         `json:"token_1_symbol,omitempty"`
	Token1Name        string         `json:"token_1_name,omitempty"`
	Token1Decimals    uint8          `json:"token_1_decimals,omitempty"`
	Token1GetReserves string         `json:"token_1_get_reserves"`
	Liquidity         string         `json:"liquidity"`
	SizeUp            bool           `json:"size_up"`
}

func (f *Callers) GetSwapPairList(chainId int) ([]*Swap, *utils.ServiceError) {
	chainIdStr := strconv.Itoa(chainId)

	var factoryAddress string
	var rpcEndPoint string
	if chainIdStr == f.conf.TestnetGiantChainId {
		factoryAddress = f.factoryAddress.TestNetworkAddress
		rpcEndPoint = f.conf.TestnetGiantEndpoint
	} else if chainIdStr == f.conf.GiantChainId {
		factoryAddress = f.factoryAddress.MainNetworkAddress
		rpcEndPoint = f.conf.GiantEndPoint
	} else {
		f.logger.Error("No value exists for chainId")
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("No value exists for chainId: ChainId=%s", chainIdStr),
			ErrorType:  utils.NoDataFound,
		}
		return nil, err
	}

	gtClient, err := ethclient.Dial(rpcEndPoint)
	if err != nil {
		f.logger.Error("gmmt json rpc error")
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("gmmt json rpc error"),
			ErrorType:  utils.InternalError,
		}
		return nil, err
	}

	factoryInstance, err := factory.NewFactory(common.HexToAddress(factoryAddress), gtClient)
	if err != nil {
		f.logger.Error("factory instance connect error")
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("factory instance connect error"),
			ErrorType:  utils.InternalError,
		}
		return nil, err
	}

	allPairsLen, err := factoryInstance.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		f.logger.Error("get pairs len error: %v", zap.Error(err))
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("get pairs len error"),
			ErrorType:  utils.InternalError,
		}
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	i := new(big.Int).SetInt64(0)
	one := new(big.Int).SetInt64(1)

	var swapList []*Swap

	for i.Cmp(allPairsLen) < 0 {
		wg.Add(1)
		go func(index *big.Int) {
			defer wg.Done()

			info := &Swap{}

			pairAddress, err := factoryInstance.AllPairs(&bind.CallOpts{}, index)
			if err != nil {
				f.logger.Error("get all pairs error: %v", zap.Error(err))
				return
			}
			pairInstance, err := pair.NewPair(pairAddress, gtClient)
			if err != nil {
				f.logger.Error("pair instance error")
				return
			}

			token0Address, err := pairInstance.Token0(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("get token 0 address error")
				return
			}
			token1Address, err := pairInstance.Token1(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("get token 1 address error")
				return
			}
			getReserves, err := pairInstance.GetReserves(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("get reserves error")
				return
			}

			token0Instance, err := erc20.NewErc20(token0Address, gtClient)
			if err != nil {
				f.logger.Error("erc20 instance error")
				return
			}
			token1Instance, err := erc20.NewErc20(token1Address, gtClient)
			if err != nil {
				f.logger.Error("erc20 instance error")
				return
			}

			token0Symbol, err := token0Instance.Symbol(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("token 0 symbol error")
				return
			}
			token1Symbol, err := token1Instance.Symbol(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("token 1 symbol error")
				return
			}

			token0Name, err := token0Instance.Name(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("token 0 name error")
				return
			}
			token1Name, err := token1Instance.Name(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("token 1 name error")
				return
			}

			token0Decimals, err := token0Instance.Decimals(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("token 0 decimals error")
				return
			}
			token1Decimals, err := token1Instance.Decimals(&bind.CallOpts{})
			if err != nil {
				f.logger.Error("token 1 decimals error")
				return
			}

			decimals0 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token0Decimals)), nil)
			decimals1 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token1Decimals)), nil)
			reserve0InEther := new(big.Float).Quo(new(big.Float).SetInt(getReserves.Reserve0), new(big.Float).SetInt(decimals0))
			reserve1InEther := new(big.Float).Quo(new(big.Float).SetInt(getReserves.Reserve1), new(big.Float).SetInt(decimals1))
			liquidity := new(big.Float).Add(reserve0InEther, reserve1InEther)

			info.Index = int(index.Int64())
			info.PairAddress = pairAddress
			info.Token0Address = token0Address
			info.Token0Symbol = token0Symbol
			info.Token0Name = token0Name
			info.Token0Decimals = token0Decimals
			info.Token0GetReserves = getReserves.Reserve0.String()
			info.Token1Address = token1Address
			info.Token1Symbol = token1Symbol
			info.Token1Name = token1Name
			info.Token1Decimals = token1Decimals
			info.Token1GetReserves = getReserves.Reserve1.String()
			info.Liquidity = liquidity.String()
			info.SizeUp = false

			mu.Lock()

			swapList = append(swapList, info)

			mu.Unlock()

		}(new(big.Int).Set(i))

		i.Add(i, one)
	}

	wg.Wait()

	sort.Slice(swapList, func(i, j int) bool {
		return swapList[i].Index < swapList[j].Index
	})

	return swapList, nil
}
