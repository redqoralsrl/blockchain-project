package service

import (
	"blockscan-go/internal/core/domain/log"
	"blockscan-go/internal/core/domain/log/log_errors"
	"blockscan-go/internal/database/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"testing"
)

type mockLogRepository struct {
	mock.Mock
}

func (m *mockLogRepository) Create(queryRower postgresql.Query, input *log.CreateLogInput) (*log.Log, error) {
	args := m.Called(queryRower, input)

	var data *log.Log
	if args.Get(0) != nil {
		data = args.Get(0).(*log.Log)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return data, err
}

func TestLogService_Create_Success(t *testing.T) {
	mockRepo := new(mockLogRepository)

	blockNumberInt := big.NewInt(1)
	zeroInt := big.NewInt(0)
	input := &log.CreateLogInput{
		ChainId:          898989,
		TransactionId:    1,
		Address:          "0x0000000000000000000000000000000000007005",
		BlockHash:        "0x0980299368c0127f3ee77f89f3c47bade65e4aa1c7d9df45b9222acc945ff833",
		BlockNumberHex:   "0x1",
		BlockNumber:      *blockNumberInt,
		Data:             "0x",
		LogIndex:         "0xd",
		Removed:          false,
		TransactionHash:  "0x906137627084ca8b611b67f5984ce34b0f7dba4505f960f57fc2257b5f53b3f2",
		TransactionIndex: "0x7",
		Timestamp:        1669713682,
		Function:         "",
		Type:             0,
		Dapp:             "",
		Collection:       "",
		FromAddress:      "0x0000000000000000000000000000000000000000",
		ToAddress:        "0x0000000000000000000000000000000000000000",
		Value:            *zeroInt,
		TokenId:          *zeroInt,
		Url:              "",
		Name:             "",
		Symbol:           "",
		Decimals:         0,
		Seller:           "0x0000000000000000000000000000000000000000",
		Buyer:            "0x0000000000000000000000000000000000000000",
		Volume:           *zeroInt,
		VolumeSymbol:     "",
		VolumeContract:   "",
		TradeNftVolume:   *zeroInt,
		Topics:           nil,
	}
	expectedData := &log.Log{
		ID:               1,
		ChainId:          898989,
		TransactionId:    1,
		Address:          "0x0000000000000000000000000000000000007005",
		BlockHash:        "0x0980299368c0127f3ee77f89f3c47bade65e4aa1c7d9df45b9222acc945ff833",
		BlockNumberHex:   "0x1",
		BlockNumber:      *blockNumberInt,
		Data:             "0x",
		LogIndex:         "0xd",
		Removed:          false,
		TransactionHash:  "0x906137627084ca8b611b67f5984ce34b0f7dba4505f960f57fc2257b5f53b3f2",
		TransactionIndex: "0x7",
		Timestamp:        1669713682,
		Function:         "",
		Type:             0,
		Dapp:             "",
		Collection:       "",
		FromAddress:      "0x0000000000000000000000000000000000000000",
		ToAddress:        "0x0000000000000000000000000000000000000000",
		Value:            *zeroInt,
		TokenId:          *zeroInt,
		Url:              "",
		Name:             "",
		Symbol:           "",
		Decimals:         0,
		Seller:           "0x0000000000000000000000000000000000000000",
		Buyer:            "0x0000000000000000000000000000000000000000",
		Volume:           *zeroInt,
		VolumeSymbol:     "",
		VolumeContract:   "",
		TradeNftVolume:   *zeroInt,
		Topics:           nil,
	}

	mockRepo.On("Create", nil, input).Return(expectedData, nil)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	result, err := service.Create(nil, input)

	assert.Nil(t, err)
	assert.Equal(t, expectedData, result)

	mockRepo.AssertExpectations(t)
}

func TestLogService_Create_Failed(t *testing.T) {
	mockRepo := new(mockLogRepository)

	blockNumberInt := big.NewInt(1)
	zeroInt := big.NewInt(0)
	input := &log.CreateLogInput{
		ChainId:          898989,
		TransactionId:    1,
		Address:          "0x0000000000000000000000000000000000007005",
		BlockHash:        "0x0980299368c0127f3ee77f89f3c47bade65e4aa1c7d9df45b9222acc945ff833",
		BlockNumberHex:   "0x1",
		BlockNumber:      *blockNumberInt,
		Data:             "0x",
		LogIndex:         "0xd",
		Removed:          false,
		TransactionHash:  "0x906137627084ca8b611b67f5984ce34b0f7dba4505f960f57fc2257b5f53b3f2",
		TransactionIndex: "0x7",
		Timestamp:        1669713682,
		Function:         "",
		Type:             0,
		Dapp:             "",
		Collection:       "",
		FromAddress:      "0x0000000000000000000000000000000000000000",
		ToAddress:        "0x0000000000000000000000000000000000000000",
		Value:            *zeroInt,
		TokenId:          *zeroInt,
		Url:              "",
		Name:             "",
		Symbol:           "",
		Decimals:         0,
		Seller:           "0x0000000000000000000000000000000000000000",
		Buyer:            "0x0000000000000000000000000000000000000000",
		Volume:           *zeroInt,
		VolumeSymbol:     "",
		VolumeContract:   "",
		TradeNftVolume:   *zeroInt,
		Topics:           nil,
	}
	expectedError := &log_errors.LogCreateError{
		StatusCode:    http.StatusInternalServerError,
		ChainId:       898989,
		BlockHash:     "0x0980299368c0127f3ee77f89f3c47bade65e4aa1c7d9df45b9222acc945ff833",
		TransactionId: 1,
		LogIndex:      "0xd",
		ErrorType:     log_errors.CreateFailedError,
	}

	mockRepo.On("Create", nil, input).Return(nil, expectedError)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	resultNil, err := service.Create(nil, input)

	assert.Nil(t, resultNil)
	assert.NotNil(t, t, err)
	assert.Equal(t, expectedError.StatusCode, err.StatusCode)
	assert.Equal(t, expectedError.ChainId, err.ChainId)
	assert.Equal(t, expectedError.BlockHash, err.BlockHash)
	assert.Equal(t, expectedError.TransactionId, err.TransactionId)
	assert.Equal(t, expectedError.LogIndex, err.LogIndex)
	assert.Equal(t, expectedError.Error(), err.Error())

	mockRepo.AssertExpectations(t)
}
