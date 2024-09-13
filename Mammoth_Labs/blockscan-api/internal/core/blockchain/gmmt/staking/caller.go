package staking

import (
	"blockscan-go/internal/config"
	"blockscan-go/internal/core/common/abi/staking"
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
	conf        *config.Config
	contractAbi *abi.ABI
	logger      *zap.Logger
}

func NewCallers(conf *config.Config, ca *abi.ABI, logger *zap.Logger) *Callers {
	return &Callers{
		conf:        conf,
		contractAbi: ca,
		logger:      logger,
	}
}

type Staking struct {
	Apr        string       `json:"apr"`
	BondedCoin string       `json:"bonded_coin"`
	List       []*Validator `json:"list"`
}

type Validator struct {
	Address        common.Address `json:"address"`
	Status         uint8          `json:"status"`
	TotalDelegated string         `json:"total_delegated"`
	SlashesCount   uint32         `json:"slashes_count"`
	ChangedAt      uint64         `json:"changed_at"`
	JailedBefore   uint64         `json:"jailed_before"`
	ClaimedAt      uint64         `json:"claimed_at"`
	CommissionRate uint16         `json:"commission_rate"`
	TotalRewards   string         `json:"total_rewards"`
}

type ValidatorByAccountData struct {
	ValidatorAddress common.Address `json:"validator_address"`
	Status           uint8          `json:"status"`
	TotalDelegated   string         `json:"total_delegated"`
	CommissionRate   uint16         `json:"commission_rate"`
	OwnerAddress     common.Address `json:"owner_address"`
	SlashesCount     uint32         `json:"slashes_count"`
	JailedBefore     uint64         `json:"jailed_before"`
	Delegation       string         `json:"delegation"`
	ClaimableAsset   string         `json:"claimable_asset"`
	PendingAsset     string         `json:"pending_asset"`
}

func (f *Callers) GetStakingList(chainId int) (*Staking, *utils.ServiceError) {
	chainIdStr := strconv.Itoa(chainId)

	var rpcEndPoint string
	if chainIdStr == f.conf.TestnetGiantChainId {
		rpcEndPoint = f.conf.TestnetGiantEndpoint
	} else if chainIdStr == f.conf.GiantChainId {
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
	stakingInstance, _ := staking.NewStaking(common.HexToAddress("0x0000000000000000000000000000000000001000"), gtClient)

	// get reward
	ecosystem, err := stakingInstance.ECOSYSTEMREWARD(&bind.CallOpts{})
	if err != nil {
		f.logger.Error("get echo system reward error: %v", zap.Error(err))
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("get echo system reward error"),
			ErrorType:  utils.InternalError,
		}
		return nil, err
	}

	// get validator list
	validators, err := stakingInstance.GetValidators(&bind.CallOpts{})
	if err != nil {
		f.logger.Error("get validators error: %v", zap.Error(err))
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("get validators error"),
			ErrorType:  utils.InternalError,
		}
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	stakingInfo := &Staking{}

	validatorList := make([]*Validator, len(validators))

	totalAmount := big.NewInt(0)

	// get validator data
	for index, validator := range validators {
		wg.Add(1)
		go func(i int, v common.Address) {
			defer wg.Done()

			validatorInfo := &Validator{}

			// check validator is jail or active
			data, err := stakingInstance.GetValidatorStatus(&bind.CallOpts{}, v)
			if err != nil {
				f.logger.Error("get validator status error: %v", zap.Error(err))
				return
			} else if data.Status == 1 {
				validatorInfo.Address = v
				validatorInfo.Status = data.Status
				validatorInfo.TotalDelegated = data.TotalDelegated.String()
				validatorInfo.SlashesCount = data.SlashesCount
				validatorInfo.ChangedAt = data.ChangedAt
				validatorInfo.JailedBefore = data.JailedBefore
				validatorInfo.ClaimedAt = data.ClaimedAt
				validatorInfo.CommissionRate = data.CommissionRate
				validatorInfo.TotalRewards = data.TotalRewards.String()

				mu.Lock()
				validatorList[i] = validatorInfo
				totalAmount.Add(totalAmount, data.TotalDelegated)
				mu.Unlock()
			} else {
				f.logger.Error("validation is jain or inactive: %v", zap.String("validate address", v.String()))
				return
			}
		}(index, validator)
	}

	wg.Wait()

	// ecosystemYear = ecosystem * 365
	ecosystemYear := new(big.Int).Mul(ecosystem, big.NewInt(365))

	// calculate APR
	ecosystemYearFloat := new(big.Float).SetInt(ecosystemYear)
	totalAmountFloat := new(big.Float).SetInt(totalAmount)
	aprFloat := new(big.Float).Quo(ecosystemYearFloat, totalAmountFloat)
	aprFloat.Mul(aprFloat, big.NewFloat(100))

	// format APR to 2 decimal places
	aprStr := aprFloat.Text('f', 2)

	stakingInfo.Apr = aprStr
	stakingInfo.BondedCoin = totalAmount.String()

	sort.Slice(validatorList, func(i, j int) bool {
		totalDelegatedI, _ := new(big.Int).SetString(validatorList[i].TotalDelegated, 10)
		totalDelegatedJ, _ := new(big.Int).SetString(validatorList[j].TotalDelegated, 10)
		return totalDelegatedI.Cmp(totalDelegatedJ) > 0
	})

	stakingInfo.List = validatorList

	return stakingInfo, nil
}

func (f *Callers) GetStakingByAccount(walletAddress string, chainId int) ([]*ValidatorByAccountData, *utils.ServiceError) {

	address := common.HexToAddress(walletAddress)
	chainIdStr := strconv.Itoa(chainId)

	var rpcEndPoint string
	if chainIdStr == f.conf.TestnetGiantChainId {
		rpcEndPoint = f.conf.TestnetGiantEndpoint
	} else if chainIdStr == f.conf.GiantChainId {
		rpcEndPoint = f.conf.GiantEndPoint
	} else {
		f.logger.Error("No value exists for chainId")
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("No value exists for chainId: ChainId=%s", chainIdStr),
			ErrorType:  utils.InternalError,
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
	stakingInstance, err := staking.NewStaking(common.HexToAddress("0x0000000000000000000000000000000000001000"), gtClient)

	// get validator list
	validators, err := stakingInstance.GetValidators(&bind.CallOpts{})
	if err != nil {
		f.logger.Error("get validators error: %v", zap.Error(err))
		err := &utils.ServiceError{
			StackTrace: zap.Stack("stacktrace").String,
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprint("get validators error"),
			ErrorType:  utils.InternalError,
		}
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	var validatorByAccountList []*ValidatorByAccountData

	for index, validator := range validators {
		wg.Add(1)
		go func(i int, v common.Address) {
			defer wg.Done()

			info := &ValidatorByAccountData{}

			data, err := stakingInstance.GetValidatorStatus(&bind.CallOpts{}, v)
			if err != nil {
				f.logger.Error("get validator status error: %v", zap.Error(err))
				return
			} else if data.Status == 1 {
				delegations, err := stakingInstance.GetValidatorDelegation(&bind.CallOpts{}, v, address)
				if err != nil {
					f.logger.Error("get validator delegation error: %v", zap.Error(err))
					return
				}
				zero := big.NewInt(0)

				if delegations.DelegatedAmount.Cmp(zero) > 0 {
					canRelegatedAmount, err := stakingInstance.CalcAvailableForRedelegateAmount(&bind.CallOpts{}, v, address)
					if err != nil {
						fmt.Println(v, address)
						f.logger.Error("calc available for redelgate amount error: %v", zap.Error(err))
						return
					}

					pendingDelegatedAssets, err := stakingInstance.GetPendingDelegatorFee(&bind.CallOpts{}, v, address)
					if err != nil {
						f.logger.Error("get pending delegator fee error: %v")
						return
					}

					info.ValidatorAddress = v
					info.Status = data.Status
					info.TotalDelegated = data.TotalDelegated.String()
					info.CommissionRate = data.CommissionRate
					info.OwnerAddress = data.OwnerAddress
					info.SlashesCount = data.SlashesCount
					info.JailedBefore = data.JailedBefore
					info.Delegation = delegations.DelegatedAmount.String()
					info.ClaimableAsset = canRelegatedAmount.AmountToStake.String()
					info.PendingAsset = new(big.Int).Sub(pendingDelegatedAssets, canRelegatedAmount.AmountToStake).String()

					mu.Lock()

					validatorByAccountList = append(validatorByAccountList, info)

					mu.Unlock()
				}
			} else {
				f.logger.Error("validation is jain or inactive: %v", zap.String("validate address", v.String()))
				return
			}

		}(index, validator)
	}

	wg.Wait()

	return validatorByAccountList, nil
}
