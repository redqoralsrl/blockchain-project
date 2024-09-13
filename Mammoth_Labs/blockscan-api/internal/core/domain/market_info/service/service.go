package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/market_info"
	"blockscan-go/internal/database/postgresql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	repo      market_info.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r market_info.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Create(queryRower postgresql.Query, input *market_info.CreateMarketInfoInput) (*market_info.MarketInfo, *utils.ServiceError) {
	marketInfoData, err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("market_info create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error market_info create: ChainId=%s LogId=%s TransactionHash=%s Collection=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId), input.TransactionHash, input.Collection),
				ErrorType:  utils.DuplicateError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("market_info create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error market_info create: ChainId=%s LogId=%s TransactionHash=%s Collection=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId), input.TransactionHash, input.Collection),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("market_info create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error market_info create: ChainId=%s LogId=%s TransactionHash=%s Collection=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId), input.TransactionHash, input.Collection),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return nil, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error market_info create: ChainId=%s LogId=%s TransactionHash=%s Collection=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId), input.TransactionHash, input.Collection),
				ErrorType:  utils.BadConnectionError,
			}
			return nil, err
		} else if strings.Contains(err.Error(), "invalid character") {
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error market_info create: ChainId=%s LogId=%s TransactionHash=%s Collection=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId), input.TransactionHash, input.Collection),
				ErrorType:  utils.JsonParseError,
			}
			return nil, err
		} else {
			s.logger.Error("market_info create error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error market_info create: ChainId=%s LogId=%s TransactionHash=%s Collection=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId), input.TransactionHash, input.Collection),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return marketInfoData, nil
}
