package service

import (
	"blockscan-go/internal/core/domain/wallet"
	"blockscan-go/internal/core/domain/wallet/wallet_errors"
	"blockscan-go/internal/database/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"testing"
)

type mockWalletRepository struct {
	mock.Mock
}

func (m *mockWalletRepository) Create(queryRower postgresql.Query, address string) (int, error) {
	args := m.Called(queryRower, address)

	var data int
	if args.Get(0) != nil {
		data = args.Get(0).(int)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return data, err
}

func (m *mockWalletRepository) Get(address string) (*wallet.Wallet, error) {
	args := m.Called(address)

	var data *wallet.Wallet
	if args.Get(0) != nil {
		data = args.Get(0).(*wallet.Wallet)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return data, err
}

func TestWalletService_Create_Success(t *testing.T) {
	mockRepo := new(mockWalletRepository)

	input := "0x85d30Cf482328965C72ae8C83e5062BdEA277Db2"
	expectedData := 1

	mockRepo.On("Create", nil, input).Return(expectedData, nil)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	result, err := service.Create(nil, input)

	assert.Nil(t, err)
	assert.Equal(t, expectedData, result)

	mockRepo.AssertExpectations(t)
}

func TestWalletService_Create_Failed(t *testing.T) {
	mockRepo := new(mockWalletRepository)

	input := "0x85d30Cf482328965C72ae8C83e5062BdEA277Db2"
	expectedData := -1
	expectedError := &wallet_errors.WalletCreateError{
		StatusCode: http.StatusInternalServerError,
		Address:    "0x85d30Cf482328965C72ae8C83e5062BdEA277Db2",
		ErrorType:  wallet_errors.CreateFailedError,
	}

	mockRepo.On("Create", nil, input).Return(expectedData, expectedError).Once()

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	result, err := service.Create(nil, input)

	assert.Equal(t, expectedData, result)
	assert.NotNil(t, t, err)
	assert.Equal(t, expectedError.Address, err.Address)
	assert.Equal(t, expectedError.StatusCode, err.StatusCode)
	assert.Equal(t, "Failed error wallet create: Address=0x85d30Cf482328965C72ae8C83e5062BdEA277Db2", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestWalletService_Get_Success(t *testing.T) {
	mockRepo := new(mockWalletRepository)

	input := "0x48157C110275391Fd8c18C981197F2D119c71C7c"
	nonce := big.NewInt(0)
	expectedData := &wallet.Wallet{
		ID:            1,
		Address:       "0x48157C110275391Fd8c18C981197F2D119c71C7c",
		NickName:      "",
		Profile:       "",
		Nonce:         *nonce,
		SignHash:      "",
		SignTimestamp: 0,
	}

	mockRepo.On("Get", input).Return(expectedData, nil)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	result, err := service.Get(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedData.ID, result.ID)
	assert.Equal(t, expectedData.Address, result.Address)
	assert.Equal(t, expectedData.Nonce, result.Nonce)

	mockRepo.AssertExpectations(t)
}

func TestWalletService_Get_Failed(t *testing.T) {
	mockRepo := new(mockWalletRepository)

	input := "0x48157C110275391Fd8c18C981197F2D119c71C7c"
	expectedError := &wallet_errors.WalletGetError{
		StatusCode: http.StatusInternalServerError,
		Address:    "0x48157C110275391Fd8c18C981197F2D119c71C7c",
		ErrorType:  wallet_errors.GetFailedError,
	}

	mockRepo.On("Get", input).Return(nil, expectedError)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	result, err := service.Get(input)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expectedError.Address, err.Address)
	assert.Equal(t, expectedError.Error(), err.Error())
	assert.Equal(t, expectedError.StatusCode, err.StatusCode)

	mockRepo.AssertExpectations(t)
}
