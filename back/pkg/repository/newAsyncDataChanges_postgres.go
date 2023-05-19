package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AsyncDataPostgres struct {
	db *sqlx.DB
}

func NewAsyncDataChanges(db *sqlx.DB) *AsyncDataPostgres {
	return &AsyncDataPostgres{db: db}
}

func (r *AsyncDataPostgres) DownSubTime() error {
	query := fmt.Sprintf("UPDATE %s SET time_of_sub=time_of_sub-1 WHERE time_of_sub<>0 AND time_of_sub<>-1", UserTable)
	_, err := r.db.Exec(query)
	return err
}
