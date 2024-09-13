package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"blockscan-go/internal/core/domain/cache"
	"time"
)

type cacheService struct {
	repo   cache.Repository
	logger *zap.Logger
}

func NewCacheService(repo cache.Repository, logger *zap.Logger) cache.UseCase {
	return &cacheService{
		repo:   repo,
		logger: logger,
	}
}

func (s *cacheService) AddOrUpdateItem(key string, value interface{}, expiration time.Duration) error {
	switch v := value.(type) {
	case string, int, int64, float64, bool:
		return s.repo.Set(key, v, expiration)
	default:
		marshaledValue, err := json.Marshal(value)
		if err != nil {
			err = &cache.CacheDataUnmarshalError{Key: key, Message: err.Error()}
			s.logger.Error("Value marshaling failed", zap.Error(err))
			return err
		}
		return s.repo.Set(key, marshaledValue, expiration)
	}
}

func (s *cacheService) RetrieveItem(key string, dest interface{}) error {
	data, err := s.repo.Get(key)
	if err != nil {
		s.logger.Error("Failed to retrieve item from cache", zap.String("key", key), zap.Error(err))
		return err
	}

	if data == nil {
		err = &cache.NoCacheDataError{Key: key}
		s.logger.Error("No data found in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	dataString, ok := data.(string)
	if !ok {
		err = &cache.CacheDataUnmarshalError{Key: key, Message: "Failed to convert cache data to string"}
		s.logger.Error("Failed to convert cache data to string", zap.String("key", key))
		return err
	}

	bytes := []byte(dataString)

	err = json.Unmarshal(bytes, dest)
	if err != nil {
		err = &cache.CacheDataUnmarshalError{Key: key, Message: err.Error()}
		s.logger.Error("Error unmarshaling item from cache", zap.String("key", key), zap.Error(err))
		return err
	}

	return nil
}

func (s *cacheService) RetrieveMultiItem(keys []string, dest interface{}) error {
	data, err := s.repo.MGet(keys)
	if err != nil {
		s.logger.Error("Failed to retrieve item from cache", zap.Strings("keys", keys), zap.Error(err))
		return err
	}

	if data == nil {
		err = &cache.NoCacheDataError{Key: keys[0]}
		s.logger.Error("No data found in cache", zap.String("key", keys[0]), zap.Error(err))
		return err
	}

	var dataString []string
	for _, v := range data {
		dataString = append(dataString, v.(string))
	}

	bytes := []byte(dataString[0])

	err = json.Unmarshal(bytes, dest)
	if err != nil {
		err = &cache.CacheDataUnmarshalError{Key: keys[0], Message: err.Error()}
		s.logger.Error("Error unmarshaling item from cache", zap.String("key", keys[0]), zap.Error(err))
		return err
	}

	return nil
}

func (s *cacheService) InvalidateItem(key string) error {
	return s.repo.Delete(key)
}

func (s *cacheService) InvalidateAll() error {
	return s.repo.Flush()
}

func (s *cacheService) ExistsItem(key string) (bool, error) {
	return s.repo.Exists(key)
}

func (s *cacheService) FindKeys(key string) ([]string, error) {
	res, err := s.repo.Keys(key)
	if err != nil {
		return nil, err
	}
	return res, nil
}
