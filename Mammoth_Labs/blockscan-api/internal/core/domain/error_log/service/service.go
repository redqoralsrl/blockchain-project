package service

import (
	"go.uber.org/zap"
	"blockscan-go/internal/core/domain/error_log"
)

type Service struct {
	repo   error_log.Repository
	logger *zap.Logger
}

func NewService(r error_log.Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   r,
		logger: logger,
	}
}

func (s *Service) Create(log *error_log.ErrorLog) (int, error) {
	num, err := s.repo.Create(log)
	if err != nil {
		return 0, err
	}

	s.logger.Info("error_log created")

	return num, err
}
