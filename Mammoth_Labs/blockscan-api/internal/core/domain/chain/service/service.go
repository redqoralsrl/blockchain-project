package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/chain"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Service struct {
	repo      chain.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r chain.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Get(chainId int) (*chain.Chain, *utils.ServiceError) {
	chainData, err := s.repo.Get(chainId)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("chain get no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error chain get: ChainId=%s", strconv.Itoa(chainId)),
				ErrorType:  utils.NotFound,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("chain get syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error chain get: ChainId=%s", strconv.Itoa(chainId)),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("chain get failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error chain get: ChainId=%s", strconv.Itoa(chainId)),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return chainData, nil
}

func (s *Service) GetToken(input *chain.GetTokenChainInput) (*chain.Chain, *utils.ServiceError) {
	chainData, err := s.repo.GetToken(input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("chain get token no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error chain get token: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.ContractAddress),
				ErrorType:  utils.NotFound,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("chain get token syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error chain get token: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.ContractAddress),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("chain get token failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error chain get token: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.ContractAddress),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return chainData, nil
}

func (s *Service) GetTokens(input *chain.GetTokensChainInput) ([]*chain.Chain, *utils.ServiceError) {
	chainList, err := s.repo.GetTokens(input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("chain get tokens no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error chain get tokens: ChainId=%s", strconv.Itoa(input.ChainId)),
				ErrorType:  utils.NotFound,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("chain get tokens syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error chain get tokens: ChainId=%s", strconv.Itoa(input.ChainId)),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("chain get token failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error chain get tokens: ChainId=%s", strconv.Itoa(input.ChainId)),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return chainList, nil
}
