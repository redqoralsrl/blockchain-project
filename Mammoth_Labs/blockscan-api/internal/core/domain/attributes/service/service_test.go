package service

import (
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/core/domain/attributes/attributes_errors"
	"blockscan-go/internal/database/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

type mockAttributesRepository struct {
	mock.Mock
}

func (m *mockAttributesRepository) Create(queryRower postgresql.Query, input *attributes.CreateAttributesInput) error {
	args := m.Called(queryRower, input)

	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}

func TestAttributesService_Create_Success(t *testing.T) {
	mockRepo := new(mockAttributesRepository)

	input := &attributes.CreateAttributesInput{
		Erc721Id:      1,
		ImageUrl:      "",
		Name:          "",
		Description:   "",
		ExternalUrl:   "",
		AttributeList: nil,
	}

	mockRepo.On("Create", nil, input).Return(nil)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	err := service.Create(nil, input)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestAttributesService_Create_Failed(t *testing.T) {
	mockRepo := new(mockAttributesRepository)

	input := &attributes.CreateAttributesInput{
		Erc721Id:      1,
		ImageUrl:      "",
		Name:          "",
		Description:   "",
		ExternalUrl:   "",
		AttributeList: nil,
	}
	expectedError := &attributes_errors.AttributesCreateError{
		StatusCode: http.StatusInternalServerError,
		Erc721Id:   1,
		ErrorType:  attributes_errors.CreateFailedError,
	}

	mockRepo.On("Create", nil, input).Return(expectedError)

	logger := zap.NewNop()
	service := NewService(mockRepo, nil, logger)

	err := service.Create(nil, input)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError.Erc721Id, err.Erc721Id)
	assert.Equal(t, expectedError.StatusCode, err.StatusCode)
	assert.Equal(t, expectedError.Error(), err.Error())

	mockRepo.AssertExpectations(t)
}
