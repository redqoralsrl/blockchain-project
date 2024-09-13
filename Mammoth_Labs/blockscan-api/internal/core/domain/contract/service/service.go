package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/contract"
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
	repo      contract.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r contract.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Get(queryRower postgresql.Query, input *contract.GetContractInput) (*contract.Contract, *utils.ServiceError) {
	contractData, err := s.repo.Get(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("contract get no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.Hash),
				ErrorType:  utils.NoDataFound,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("contract get syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.Hash),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("contract get failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract get: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.Hash),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return contractData, nil
}

func (s *Service) Create(queryRower postgresql.Query, input *contract.CreateContractInput) (*contract.Contract, *utils.ServiceError) {
	contractData, err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			getContract := &contract.GetContractInput{
				ChainId: input.ChainId,
				Hash:    input.Hash,
			}
			data, getErr := s.Get(queryRower, getContract)
			if getErr != nil {
				s.logger.Error("contract create error duplicate", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusConflict,
					Message:    fmt.Sprintf("error contract create: ChainId=%s TransactionId=%s ContractHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.TransactionCreateId), input.Hash),
					ErrorType:  utils.DuplicateError,
				}
				return nil, err
			}
			return data, nil
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("contract create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract create: ChainId=%s TransactionId=%s ContractHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.TransactionCreateId), input.Hash),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("contract create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract create: ChainId=%s TransactionId=%s ContractHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.TransactionCreateId), input.Hash),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return nil, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("contract create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error contract create: ChainId=%s TransactionId=%s ContractHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.TransactionCreateId), input.Hash),
				ErrorType:  utils.BadConnectionError,
			}
			return nil, err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("contract create error Json parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract create: ChainId=%s TransactionId=%s ContractHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.TransactionCreateId), input.Hash),
				ErrorType:  utils.JsonParseError,
			}
			return nil, err
		} else {
			s.logger.Error("contract create failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract create: ChainId=%s TransactionId=%s ContractHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.TransactionCreateId), input.Hash),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return contractData, nil
}

func (s *Service) GetId(queryRower postgresql.Query, input *contract.GetContractIdInput) (int, *utils.ServiceError) {
	data, err := s.repo.GetId(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("contract get id no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get id: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.ContractAddress),
				ErrorType:  utils.NoDataFound,
			}
			return data, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("contract get id syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get id: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.ContractAddress),
				ErrorType:  utils.SyntaxError,
			}
			return data, err
		} else {
			s.logger.Error("contract get type failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract get id: ChainId=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.ContractAddress),
				ErrorType:  utils.Fail,
			}
			return data, err
		}

	}

	return data, nil
}

func (s *Service) GetType(queryRower postgresql.Query, input *contract.GetContractTypeInput) (int, *utils.ServiceError) {
	data, err := s.repo.GetType(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Error("contract get type no_data_found", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get type: ChainId=%s Hash=%s", strconv.Itoa(input.ChainId), input.Hash),
				ErrorType:  utils.NoDataFound,
			}
			return data, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("contract get type syntax_error", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get type: ChainId=%s Hash=%s", strconv.Itoa(input.ChainId), input.Hash),
				ErrorType:  utils.SyntaxError,
			}
			return data, err
		} else {
			s.logger.Error("contract get type failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract get type: ChainId=%s Hash=%s", strconv.Itoa(input.ChainId), input.Hash),
				ErrorType:  utils.Fail,
			}
			return data, err
		}
	}

	return data, nil
}

func (s *Service) UpdateType(queryRower postgresql.Query, input *contract.UpdateContractTypeInput) *utils.ServiceError {
	err := s.repo.UpdateType(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("contract update type error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error contract update type: ChainId=%s Hash=%s Type=%s", strconv.Itoa(input.ChainId), input.Hash, strconv.Itoa(input.Type)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "foreign_key_violation" {
			s.logger.Error("contract update type error foreign_key", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract update type: ChainId=%s Hash=%s Type=%s", strconv.Itoa(input.ChainId), input.Hash, strconv.Itoa(input.Type)),
				ErrorType:  utils.ForeignKeyError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "check_violation" {
			s.logger.Error("contract update type error check_violation", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract update type: ChainId=%s Hash=%s Type=%s", strconv.Itoa(input.ChainId), input.Hash, strconv.Itoa(input.Type)),
				ErrorType:  utils.CheckViolationError,
			}
			return err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("contract update type error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error contract update type: ChainId=%s Hash=%s Type=%s", strconv.Itoa(input.ChainId), input.Hash, strconv.Itoa(input.Type)),
				ErrorType:  utils.BadConnectionError,
			}
			return err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("contract update type error invalid character", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract update type: ChainId=%s Hash=%s Type=%s", strconv.Itoa(input.ChainId), input.Hash, strconv.Itoa(input.Type)),
				ErrorType:  utils.JsonParseError,
			}
			return err
		} else {
			s.logger.Error("contract update type error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract update type: ChainId=%s Hash=%s Type=%s", strconv.Itoa(input.ChainId), input.Hash, strconv.Itoa(input.Type)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) UpdateAllVolume(queryRower postgresql.Query, input *contract.UpdateContractAllVolumeInput) *utils.ServiceError {
	err := s.repo.UpdateAllVolume(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("contract update all volume error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error contract update all volume: ChainId=%s Hash=%s Volume=%s", strconv.Itoa(input.ChainId), input.Hash, input.Volume.String()),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "foreign_key_violation" {
			s.logger.Error("contract update all volume error foreign_key", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract update all volume: ChainId=%s Hash=%s Volume=%s", strconv.Itoa(input.ChainId), input.Hash, input.Volume.String()),
				ErrorType:  utils.ForeignKeyError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "check_violation" {
			s.logger.Error("contract update all volume error check_violation", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract update all volume: ChainId=%s Hash=%s Volume=%s", strconv.Itoa(input.ChainId), input.Hash, input.Volume.String()),
				ErrorType:  utils.CheckViolationError,
			}
			return err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("contract update all volume error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error contract update all volume: ChainId=%s Hash=%s Volume=%s", strconv.Itoa(input.ChainId), input.Hash, input.Volume.String()),
				ErrorType:  utils.BadConnectionError,
			}
			return err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("contract update all volume error Json parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract update all volume: ChainId=%s Hash=%s Volume=%s", strconv.Itoa(input.ChainId), input.Hash, input.Volume.String()),
				ErrorType:  utils.JsonParseError,
			}
			return err
		} else {
			s.logger.Error("contract update all volume error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract update all volume: ChainId=%s Hash=%s Volume=%s", strconv.Itoa(input.ChainId), input.Hash, input.Volume.String()),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) GetWalletNFTsByCollection(input *contract.GetContractWalletNFTsByCollectionInput) ([]*contract.GetContractWalletNFTsByCollectionData, *utils.ServiceError) {
	data, err := s.repo.GetWalletNFTsByCollection(input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			return []*contract.GetContractWalletNFTsByCollectionData{}, nil
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("contract get wallet nfts by collection syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get wallet nfts by collection: ChainId=%s WalletAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("contract get wallet nfts by collection failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract get wallet nfts by collection: ChainId=%s WalletAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return data, nil
}

func (s *Service) GetCollectionNFTsForWallet(input *contract.GetContractCollectionNFTsForWalletInput) (*contract.GetContractCollectionNFTsForWalletData, *utils.ServiceError) {
	data, err := s.repo.GetCollectionNFTsForWallet(input)

	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("contract get collection NFTs for wallet syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error contract get collection NFTs for wallet: WalletAddress=%s Hash=%s ChainId=%s", input.WalletAddress, input.Hash, strconv.Itoa(input.ChainId)),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("contract get collection NFTs for wallet failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error contract  get collection NFTs for wallet: WalletAddress=%s Hash=%s ChainId=%s", input.WalletAddress, input.Hash, strconv.Itoa(input.ChainId)),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return data, nil
}
