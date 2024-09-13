package service

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/nft_volume"
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
	repo      nft_volume.Repository
	txManager postgresql.DBTransactionManager
	logger    *zap.Logger
}

func NewService(r nft_volume.Repository, tx postgresql.DBTransactionManager, logger *zap.Logger) *Service {
	return &Service{
		repo:      r,
		txManager: tx,
		logger:    logger,
	}
}

func (s *Service) Create(queryRower postgresql.Query, input *nft_volume.CreateNftVolumeInput) (*nft_volume.NftVolume, *utils.ServiceError) {
	nftVolumeData, err := s.repo.Create(queryRower, input)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			s.logger.Error("nft_volume create error duplicate", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusConflict,
				Message:    fmt.Sprintf("error nft_volume create: ChainId=%s LogId=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId)),
				ErrorType:  utils.DuplicateError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "syntax_error" {
			s.logger.Error("nft_volume create error syntax", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error nft_volume create: ChainId=%s LogId=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId)),
				ErrorType:  utils.SyntaxError,
			}
			return nil, err
		} else if errors.As(err, &pqErr) && pqErr.Code.Name() == "invalid_sql_statement_name" {
			s.logger.Error("nft_volume create error invalid_sql_statement_name", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error nft_volume create: ChainId=%s LogId=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId)),
				ErrorType:  utils.InvalidSQLStatementError,
			}
			return nil, err
		} else if err.Error() == "driver: bad connection" || err.Error() == "pq: unexpected Parse response 'D'" || err.Error() == "pq: unexpected Parse response 'C'" {
			//s.logger.Error("nft_volume create error bad connection", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusServiceUnavailable,
				Message:    fmt.Sprintf("error nft_volume create: ChainId=%s LogId=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId)),
				ErrorType:  utils.BadConnectionError,
			}
			return nil, err
		} else if strings.Contains(err.Error(), "invalid character") {
			//s.logger.Error("nft_volume create error Json parsing", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusBadRequest,
				Message:    fmt.Sprintf("error nft_volume create: ChainId=%s LogId=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId)),
				ErrorType:  utils.JsonParseError,
			}
			return nil, err
		} else {
			s.logger.Error("nft_volume create failed", zap.Error(err))
			err := &utils.ServiceError{
				StackTrace: zap.Stack("stacktrace").String,
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("error nft_volume create: ChainId=%s LogId=%s", strconv.Itoa(input.ChainId), strconv.Itoa(input.LogId)),
				ErrorType:  utils.Fail,
			}
			return nil, err
		}
	}

	return nftVolumeData, nil
}
