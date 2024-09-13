package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/nft"
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
	repo      nft.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r nft.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Get(input *nft.GetNftInput) ([]*nft.GetNftData, *utils.ServiceError) {
	nftList, err := s.repo.Get(input)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*nft.GetNftData{}, nil
		} else {
			var pqErr *pq.Error

			if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("nft get syntax_error", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error nft get syntax_error: WalletAddress=%s ChainId=%s", input.WalletAddress, strconv.Itoa(input.ChainId)),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("nft get failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error nft get failed: WalletAddress=%s ChainId=%s", input.WalletAddress, strconv.Itoa(input.ChainId)),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
	}

	return nftList, nil
}

func (s *Service) GetDetail(input *nft.GetNftDetailInput) (*nft.GetNFtDetailData, *utils.ServiceError) {
	nftData, err := s.repo.GetDetail(input)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			var pqErr *pq.Error

			if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("nft get detail syntax_error", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error nft get detail syntax_error: ChainId=%s Hash=%s TokenId=%s", strconv.Itoa(input.ChainId), input.Hash, input.TokenId),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("nft get detail failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error nft get detail failed: ChainId=%s Hash=%s TokenId=%s", strconv.Itoa(input.ChainId), input.Hash, input.TokenId),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
	}

	return nftData, nil
}

func (s *Service) GetDetailOfWallet(input *nft.GetNftDetailOfWalletInput) (*nft.GetNFtDetailData, *utils.ServiceError) {
	nftData, err := s.repo.GetDetailOfWallet(input)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			var pqErr *pq.Error

			if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("nft get detail of wallet syntax_error", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error nft get detail of wallet syntax_error: WalletAddress=%s ChainId=%s Hash=%s TokenId=%s", input.WalletAddress, strconv.Itoa(input.ChainId), input.Hash, input.TokenId),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("nft get detail of wallet failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error nft get detail of wallet failed: WalletAddress=%s ChainId=%s Hash=%s TokenId=%s", input.WalletAddress, strconv.Itoa(input.ChainId), input.Hash, input.TokenId),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
	}

	return nftData, nil
}
