package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/core/domain/erc1155"
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
	repo      erc1155.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r erc1155.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) MoveErc1155(queryRower postgresql.Query, input *erc1155.MoveErc1155Input) *utils.ServiceError {
	err := s.repo.MoveErc1155(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("erc1155 move error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error erc1155 move: ChainId=%s ContractId=%s From=%s To=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.ContractId), input.From, input.To),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("erc1155 move error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error erc1155 move: ChainId=%s ContractId=%s From=%s To=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.ContractId), input.From, input.To),
				ErrorType:  utils.SyntaxError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("erc1155 move error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error erc1155 move: ChainId=%s ContractId=%s From=%s To=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.ContractId), input.From, input.To),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" || err.Error() == "pq: unexpected Parse response 'C'" {
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error erc1155 move: ChainId=%s ContractId=%s From=%s To=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.ContractId), input.From, input.To),
				ErrorType:  utils.BadConnectionError,
			}
			return err
		} else if strings.Contains(err.Error(), "invalid character") {
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error erc1155 move: ChainId=%s ContractId=%s From=%s To=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.ContractId), input.From, input.To),
				ErrorType:  utils.JsonParseError,
			}
			return err
		} else {
			s.logger.Error("erc1155 move error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error erc1155 move: ChainId=%s ContractId=%s From=%s To=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.ContractId), input.From, input.To),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) GetEmptyUrlErc1155List() []*erc1155.GetEmptyUrlErc1155 {
	ercList, err := s.repo.GetEmptyUrlErc1155List()

	if err != nil {
		return []*erc1155.GetEmptyUrlErc1155{}
	}

	return ercList
}

func (s *Service) UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) *utils.ServiceError {
	err := s.repo.UpdateIsUndefinedMetaData(queryRower, id)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("erc1155 update isUndefinedMetaData error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error erc1155 update is undefined meta data: Erc1155Id=%s", strconv.Itoa(id)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "foreign_key_violation" {
			s.logger.Error("erc1155 update isUndefinedMetaData error foreign_key", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error erc1155 update is undefined meta data: Erc1155Id=%s", strconv.Itoa(id)),
				ErrorType:  utils.ForeignKeyError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "check_violation" {
			s.logger.Error("erc1155 update isUndefinedMetaData error check_violation", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error erc1155 update is undefined meta data: Erc1155Id=%s", strconv.Itoa(id)),
				ErrorType:  utils.CheckViolationError,
			}
			return err
		} else {
			s.logger.Error("erc1155 update isUndefinedMetaData error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error erc1155 update is undefined meta data: Erc1155Id=%s", strconv.Itoa(id)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) Update(queryRower postgresql.Query, input *attributes.CreateErc1155AttributesInput) *utils.ServiceError {
	err := s.repo.Update(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("erc1155 update error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error erc721 update: Erc1155Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "foreign_key_violation" {
			s.logger.Error("erc1155 update error foreign_key", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error erc721 update: Erc1155Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.ForeignKeyError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "check_violation" {
			s.logger.Error("erc1155 update error check_violation", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error erc721 update: Erc1155Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.CheckViolationError,
			}
			return err
		} else {
			s.logger.Error("erc1155 update error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error erc721 update: Erc1155Id=%s", strconv.Itoa(input.Erc1155Id)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}
