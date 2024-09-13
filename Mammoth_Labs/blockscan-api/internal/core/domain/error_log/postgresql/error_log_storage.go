package postgresql

import (
	"database/sql"
	"blockscan-go/internal/core/domain/error_log"
)

type ErrorLogStorage struct {
	db *sql.DB
}

func NewErrorLogStorage(db *sql.DB) *ErrorLogStorage {
	return &ErrorLogStorage{db}
}

func (s *ErrorLogStorage) Create(log *error_log.ErrorLog) (int, error) {
	if log.Timestamp.IsZero() {
		log.Timestamp = log.Timestamp.UTC()
	}

	query := `
		insert into error_log (timestamp, ip_address, user_agent, path, http_method, requested_url, error_code, error_message, stack_trace) 
		values ($1, $2, $3, $4, $5,$6, $7, $8, $9)
		returning id;
	`

	var id int
	err := s.db.QueryRow(
		query,
		log.Timestamp,
		log.IPAddress,
		log.UserAgent,
		log.Path,
		log.HttpMethod,
		log.RequestUrl,
		log.ErrorCode,
		log.ErrorMessage,
		log.StackTrace,
	).Scan(
		&id,
	)

	return id, err
}
