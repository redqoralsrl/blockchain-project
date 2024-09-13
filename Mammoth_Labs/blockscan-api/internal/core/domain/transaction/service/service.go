package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/transaction"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	repo      transaction.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r transaction.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Create(queryRower postgresql.Query, input *transaction.CreateTransactionInput) (*transaction.Transaction, *utils.ServiceError) {
	transactionData, err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("transaction create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error transaction create: ChainId=%s BlockNumber=%s BlockHash=%s TransactionHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.BlockNumber.Uint64())), input.BlockHash, input.Hash),
				ErrorType:  utils.DuplicateError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("transaction create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error transaction create: ChainId=%s BlockNumber=%s BlockHash=%s TransactionHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.BlockNumber.Uint64())), input.BlockHash, input.Hash),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("transaction create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error transaction create: ChainId=%s BlockNumber=%s BlockHash=%s TransactionHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.BlockNumber.Uint64())), input.BlockHash, input.Hash),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return nil, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("transaction create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error transaction create: ChainId=%s BlockNumber=%s BlockHash=%s TransactionHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.BlockNumber.Uint64())), input.BlockHash, input.Hash),
				ErrorType:  utils.BadConnectionError,
			}
			return nil, err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("transaction create error JSON parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error transaction create: ChainId=%s BlockNumber=%s BlockHash=%s TransactionHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.BlockNumber.Uint64())), input.BlockHash, input.Hash),
				ErrorType:  utils.JsonParseError,
			}
			return nil, err
		} else {
			s.logger.Error("transaction create failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error transaction create: ChainId=%s BlockNumber=%s BlockHash=%s TransactionHash=%s", strconv.Itoa(input.ChainId), strconv.Itoa(int(input.BlockNumber.Uint64())), input.BlockHash, input.Hash),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return transactionData, nil
}

func (s *Service) Update(queryRower postgresql.Query, input *transaction.UpdateTransactionInput) *utils.ServiceError {
	err := s.repo.Update(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("transaction update error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error transaction update: TransactionId=%s ContractId=%s", strconv.Itoa(input.TransactionId), strconv.Itoa(input.ContractId)),
				ErrorType:  utils.DuplicateError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "foreign_key_violation" {
			s.logger.Error("transaction update error foreign_key", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error transaction update: TransactionId=%s ContractId=%s", strconv.Itoa(input.TransactionId), strconv.Itoa(input.ContractId)),
				ErrorType:  utils.ForeignKeyError,
			}
			return err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "check_violation" {
			s.logger.Error("transaction update error check_violation", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error transaction update: TransactionId=%s ContractId=%s", strconv.Itoa(input.TransactionId), strconv.Itoa(input.ContractId)),
				ErrorType:  utils.CheckViolationError,
			}
			return err
		} else {
			s.logger.Error("transaction update error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error transaction update: TransactionId=%s ContractId=%s", strconv.Itoa(input.TransactionId), strconv.Itoa(input.ContractId)),
				ErrorType:  utils.Fail,
			}
			return err
		}
	}

	return nil
}

func (s *Service) Get(input *transaction.GetTransactionInput) ([]*transaction.GetTransactionData, *utils.ServiceError) {
	var transactions []*transaction.GetTransactionData
	if strings.EqualFold(input.ChainType, "token") {
		tx, err := s.repo.GetToken(input)
		if err != nil {
			var pqErr *pq.Error

			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("transaction get token error syntax", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error get token transaction: ChainId=%s WalletAddress=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress, input.ContractAddress),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("transaction get token error failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error get token transaction: ChainId=%s WalletAddress=%s ContractAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress, input.ContractAddress),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
		transactions = tx
	} else {
		tx, err := s.repo.GetCoin(input)
		if err != nil {
			var pqErr *pq.Error

			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
				s.logger.Error("transaction get coin error syntax", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusBadRequest,
					Message:    fmt.Sprintf("error get coin transaction: ChainId=%s WalletAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress),
					ErrorType:  utils.SyntaxError,
				}
				return nil, err
			} else {
				s.logger.Error("transaction get coin error failed", zap.Error(err))
				err := &utils.ServiceError{
					StackTrace: zap.Stack("stacktrace").String,
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("error get coin transaction: ChainId=%s WalletAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress),
					ErrorType:  utils.Fail,
				}
				return nil, err
			}
		}
		transactions = tx
	}

	calculatedTransactions := make([]*transaction.GetTransactionData, len(transactions))
	for index, tx := range transactions {
		calculateTxFee, err := calculateGasFee(tx)
		if err != nil {
			s.logger.Error("Error calculate transactionFee", zap.Error(err))
			continue
		}
		calculatedTransactions[index] = calculateTxFee
	}

	return calculatedTransactions, nil
}

func (s *Service) GetAll(input *transaction.GetAllTransactionInput) ([]*transaction.GetTransactionData, *utils.ServiceError) {
	transactions, err := s.repo.GetAll(input)
	if err != nil {
		var pqErr *pq.Error

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("transaction get all error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error get all transaction: ChainId=%s WalletAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else {
			s.logger.Error("transaction get all error failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error get all transaction: ChainId=%s WalletAddress=%s", strconv.Itoa(input.ChainId), input.WalletAddress),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	calculatedTransactions := make([]*transaction.GetTransactionData, len(transactions))
	for index, tx := range transactions {
		calculateTxFee, err := calculateGasFee(tx)
		if err != nil {
			s.logger.Error("Error calculate transactionFee", zap.Error(err))
			continue
		}
		calculatedTransactions[index] = calculateTxFee
	}

	return calculatedTransactions, nil
}

func calculateGasFee(tx *transaction.GetTransactionData) (*transaction.GetTransactionData, error) {
	if !strings.HasPrefix(tx.Gas, "0x") {
		return nil, fmt.Errorf("gas value must hasve '0x' prefix")
	}
	if !strings.HasPrefix(tx.GasPrice, "0x") {
		return nil, fmt.Errorf("gas price value must hasve '0x' prefix")
	}

	gas, ok := new(big.Int).SetString(tx.Gas[2:], 16)
	if !ok {
		return nil, fmt.Errorf("failed to parse gas")
	}
	gasPrice, ok := new(big.Int).SetString(tx.GasPrice[2:], 16)
	if !ok {
		return nil, fmt.Errorf("failed to parse gas price")
	}
	txFee := new(big.Int).Mul(gas, gasPrice)

	calculatedGasUsed := *tx
	calculatedGasUsed.TransactionFee = "0x" + txFee.Text(16)

	return &calculatedGasUsed, nil
}
