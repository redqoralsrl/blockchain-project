package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/log"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	repo      log.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r log.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Create(queryRower postgresql.Query, input *log.CreateLogInput) (*log.Log, *utils.ServiceError) {
	logData, err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("log create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error log create: ChainId=%s BlockHash=%s TransactionId=%s LogIndex=%s", strconv.Itoa(input.ChainId), input.BlockHash, strconv.Itoa(input.TransactionId), input.LogIndex),
				ErrorType:  utils.DuplicateError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("log create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error log create: ChainId=%s BlockHash=%s TransactionId=%s LogIndex=%s", strconv.Itoa(input.ChainId), input.BlockHash, strconv.Itoa(input.TransactionId), input.LogIndex),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("log create error invalid_sql_statement_name")
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error log create: ChainId=%s BlockHash=%s TransactionId=%s LogIndex=%s", strconv.Itoa(input.ChainId), input.BlockHash, strconv.Itoa(input.TransactionId), input.LogIndex),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return nil, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("log create error bad connection")
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error log create: ChainId=%s BlockHash=%s TransactionId=%s LogIndex=%s", strconv.Itoa(input.ChainId), input.BlockHash, strconv.Itoa(input.TransactionId), input.LogIndex),
				ErrorType:  utils.BadConnectionError,
			}
			return nil, err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("log create error Json parsing")
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error log create: ChainId=%s BlockHash=%s TransactionId=%s LogIndex=%s", strconv.Itoa(input.ChainId), input.BlockHash, strconv.Itoa(input.TransactionId), input.LogIndex),
				ErrorType:  utils.JsonParseError,
			}
			return nil, err
		} else {
			s.logger.Error("log create falied", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error log create: ChainId=%s BlockHash=%s TransactionId=%s LogIndex=%s", strconv.Itoa(input.ChainId), input.BlockHash, strconv.Itoa(input.TransactionId), input.LogIndex),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return logData, nil
}

func (s *Service) GetNFTsByWallet(input *log.GetLogNFTsByWalletInput) ([]*log.GetLogNFTsByWalletData, *utils.ServiceError) {
	logs, err := s.repo.GetNFTsByWallet(input)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*log.GetLogNFTsByWalletData{}, nil
		} else {
			var pqErr *pq.Error

			if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("log get NFTs by wallet syntax_error", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error log get NFTs by wallet: WalletAddress=%s", input.WalletAddress),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("log get NFTs by wallet failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error log get NFTs by wallet: WalletAddress=%s", input.WalletAddress),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
	}

	return logs, nil
}

func (s *Service) GetSearchNFTsByWallet(input *log.GetSearchNFTsByWalletInput) ([]*log.GetLogNFTsByWalletData, *utils.ServiceError) {
	logs, err := s.repo.GetSearchNFTsByWallet(input)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*log.GetLogNFTsByWalletData{}, nil
		} else {
			var pqErr *pq.Error

			if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("log get search NFTs by wallet syntax_error", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error log get search NFTs by wallet: WalletAddress=%s", input.WalletAddress),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("log get search NFTs by wallet failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error log get search NFTs by wallet: WalletAddress=%s", input.WalletAddress),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
	}

	return logs, nil
}
