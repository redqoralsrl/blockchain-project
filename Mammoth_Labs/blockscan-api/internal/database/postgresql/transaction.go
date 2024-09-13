package postgresql

import (
	"database/sql"
)

type DBTransactionManager interface {
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}

type Manager struct {
	db *sql.DB
}

func NewManager(db *sql.DB) *Manager {
	return &Manager{db: db}
}

func (m *Manager) Begin() (*sql.Tx, error) {
	return m.db.Begin()
}

func (m *Manager) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (m *Manager) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
