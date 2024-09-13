package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/wallet"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type Service struct {
	repo      wallet.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r wallet.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Get(queryRower postgresql.Query, address string) (*wallet.Wallet, *utils.ServiceError) {
	data, err := s.repo.Get(queryRower, address)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("wallet get error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error wallet get: Address=%s", address),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("wallet get failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error wallet get: Address=%s", address),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return data, nil
}

func (s *Service) Create(queryRower postgresql.Query, address string) (int, *utils.ServiceError) {
	walletId, err := s.repo.Create(queryRower, address)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("wallet create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error wallet create: Address=%s", address),
				ErrorType:  utils.DuplicateError,
			}
			return -1, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("wallet create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error wallet create: Address=%s", address),
				ErrorType:  utils.SyntaxError,
			}
			return -1, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("wallet create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error wallet create: Address=%s", address),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return -1, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			// s.logger.Error("wallet create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error wallet create: Address=%s", address),
				ErrorType:  utils.BadConnectionError,
			}
			return -1, err
		} else if strings.Contains(err.Error(), "invalid character") {
			// s.logger.Error("wallet create error JSON parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error wallet create: Address=%s", address),
				ErrorType:  utils.JsonParseError,
			}
			return -1, err
		} else {
			s.logger.Error("wallet create failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error wallet create: Address=%s", address),
				ErrorType:  utils.Fail,
			}
			return -1, err
		}
	}

	return walletId, nil
}
