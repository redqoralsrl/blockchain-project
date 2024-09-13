package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/tracking_info"
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
	repo      tracking_info.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r tracking_info.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Get(chainId int) (*tracking_info.TrackingInfo, *utils.ServiceError) {
	data, err := s.repo.Get(chainId)

	if err != nil {
		var pqErr *pq.Error

		chainIdString := strconv.Itoa(chainId)
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("tracking_info get no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error tracking_info get: chain_id=%s", chainIdString),
				ErrorType:  utils.NoDataFound,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("tracking_info get syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error tracking_info get: chain_id=%s", chainIdString),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("tracking_info get failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error tracking_info get: chain_id=%s", chainIdString),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return data, nil
}

func (s *Service) Increase(chainId int) *utils.ServiceError {
	err := s.repo.Increase(nil, chainId)

	if err != nil {
		var pqErr *pq.Error

		chainIdString := strconv.Itoa(chainId)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("tracking_info increase invalid_sql_statement_name ", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error tracking_info increase: chain_id=%s", chainIdString),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return err
		} else {
			s.logger.Error("tracking_info increase failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error tracking_info increase: chain_id=%s", chainIdString),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}
