package cache

import "fmt"

type NoCacheDataError struct {
	Key string
}

func (e *NoCacheDataError) Error() string {
	return fmt.Sprintf("no cache data for key: %s", e.Key)
}

type CacheDataUnmarshalError struct {
	Key     string
	Message string
}

func (e *CacheDataUnmarshalError) Error() string {
	return fmt.Sprintf("failed to unmarshal cache data for key: %s , message: %s", e.Key, e.Message)
}
