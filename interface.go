package gorm

import (
	"context"
	"database/sql"
)

// SQLCommon is the minimal database connection functionality gorm requires.  Implemented by *sql.DB.
type SQLCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row 
}

type sqlDb interface {
	Begin() (*sql.Tx, error)
}

type sqlTx interface {
	Commit() error
	Rollback() error
}
