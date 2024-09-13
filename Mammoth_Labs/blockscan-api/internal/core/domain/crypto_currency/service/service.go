package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/crypto_currency"
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
	repo      crypto_currency.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r crypto_currency.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Create(queryRower postgresql.Query, input *crypto_currency.CreateCryptoCurrencyInput) *utils.ServiceError {
	err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("crypto_currency create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error crypto_currency create: Timestamp=%s", strconv.Itoa(input.Timestamp)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("crypto_currency create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error crypto_currency create: Timestamp=%s", strconv.Itoa(input.Timestamp)),
				ErrorType:  utils.SyntaxError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("crypto_currency create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error crypto_currency create: Timestamp=%s", strconv.Itoa(input.Timestamp)),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return err
		} else {
			s.logger.Error("crypto_currency create error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error crypto_currency create: Timestamp=%s", strconv.Itoa(input.Timestamp)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) Get() (*crypto_currency.GetCryptoCurrencyData, *utils.ServiceError) {
	data, err := s.repo.Get()

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("crypto_currency get no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error crypto currency get"),
				ErrorType:  utils.NoDataFound,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("crypto_currency get syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error crypto currency get"),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("crypto_currency get failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error crypto currency get"),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return data, nil
}
