package repository

import (
	"fmt"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user grod.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, subscribe) VALUES ($1, $2, $3, $4) RETURNING id", UserTable)

	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, "free")
	if err := row.Scan(&id); err != nil {
		return 300, err
	}
	return 200, nil
}
