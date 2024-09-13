package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/block"
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
	repo      block.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r block.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Create(queryRower postgresql.Query, input *block.CreateBlockInput) (*block.Block, *utils.ServiceError) {
	blockData, err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("block create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error block create: ChainId=%s BlockNumber=%s BlockHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.Number.Uint64())), input.Hash),
				ErrorType:  utils.DuplicateError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("block create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error block create: ChainId=%s BlockNumber=%s BlockHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.Number.Uint64())), input.Hash),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("block create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error block create: ChainId=%s BlockNumber=%s BlockHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.Number.Uint64())), input.Hash),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return nil, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("block create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error block create: ChainId=%s BlockNumber=%s BlockHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.Number.Uint64())), input.Hash),
				ErrorType:  utils.BadConnectionError,
			}
			return nil, err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("block create error Json parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error block create: ChainId=%s BlockNumber=%s BlockHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.Number.Uint64())), input.Hash),
				ErrorType:  utils.JsonParseError,
			}
			return nil, err
		} else {
			s.logger.Error("block create failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error block create: ChainId=%s BlockNumber=%s BlockHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.Number.Uint64())), input.Hash),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return blockData, nil
}
