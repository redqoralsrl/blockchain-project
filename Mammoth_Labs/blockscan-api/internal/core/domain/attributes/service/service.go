package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/attributes"
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
	repo      attributes.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r attributes.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) CreateErc721(queryRower postgresql.Query, input *attributes.CreateErc721AttributesInput) *utils.ServiceError {
	err := s.repo.CreateErc721(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("attributes create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc721Id)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("attributes create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc721Id)),
				ErrorType:  utils.SyntaxError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("attributes create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc721Id)),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("attributes create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc721Id)),
				ErrorType:  utils.BadConnectionError,
			}
			return err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("attributes create error Json parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc721Id)),
				ErrorType:  utils.JsonParseError,
			}
			return err
		} else {
			s.logger.Error("attributes create error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc721Id)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) CreateErc1155(queryRower postgresql.Query, input *attributes.CreateErc1155AttributesInput) *utils.ServiceError {
	err := s.repo.CreateErc1155(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("attributes create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("attributes create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.SyntaxError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("attributes create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("attributes create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.BadConnectionError,
			}
			return err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("attributes create error Json parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.JsonParseError,
			}
			return err
		} else {
			s.logger.Error("attributes create error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error attributes create: Erc721Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}
