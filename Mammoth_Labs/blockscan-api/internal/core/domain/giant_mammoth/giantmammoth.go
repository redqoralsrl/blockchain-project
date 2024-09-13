package giant_mammoth

import (
	"blockscan-go/internal/core/blockchain/gmmt/staking"
	"blockscan-go/internal/core/blockchain/gmmt/swap"
	"blockscan-go/internal/core/common/utils"
)

type StakingUseCase interface {
	GetStakingList(chainId int) (*staking.Staking, *utils.ServiceError)
	GetStakingByAccount(walletAddress string, chainId int) ([]*staking.ValidatorByAccountData, *utils.ServiceError)
}

type SwapUseCase interface {
	GetSwapPairList(chainId int) ([]*swap.Swap, *utils.ServiceError)
}

type UseCase interface {
	StakingUseCase
	SwapUseCase
}

type StakingAdapter interface {
	GetStakingList(chainId int) (*staking.Staking, *utils.ServiceError)
	GetStakingByAccount(walletAddress string, chainId int) ([]*staking.ValidatorByAccountData, *utils.ServiceError)
}

type SwapAdapter interface {
	GetSwapPairList(chainId int) ([]*swap.Swap, *utils.ServiceError)
}
