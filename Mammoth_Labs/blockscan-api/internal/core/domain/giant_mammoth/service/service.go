package service

import (
	"blockscan-go/internal/core/blockchain/gmmt/staking"
	"blockscan-go/internal/core/blockchain/gmmt/swap"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/giant_mammoth"
	"go.uber.org/zap"
)

type Service struct {
	staking giant_mammoth.StakingAdapter
	swap    giant_mammoth.SwapAdapter
	logger  *zap.Logger
}

func NewService(staking giant_mammoth.StakingAdapter, swap giant_mammoth.SwapAdapter, logger *zap.Logger) *Service {
	return &Service{
		staking: staking,
		swap:    swap,
		logger:  logger,
	}
}

func (s *Service) GetStakingList(chainId int) (*staking.Staking, *utils.ServiceError) {
	data, err := s.staking.GetStakingList(chainId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) GetStakingByAccount(walletAddress string, chainId int) ([]*staking.ValidatorByAccountData, *utils.ServiceError) {
	data, err := s.staking.GetStakingByAccount(walletAddress, chainId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) GetSwapPairList(chainId int) ([]*swap.Swap, *utils.ServiceError) {
	data, err := s.swap.GetSwapPairList(chainId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
