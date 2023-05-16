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
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email, subscribe, time_of_sub, count_requests) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", UserTable)

	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, "free", user.TimeSub, user.CountRequests)
	if err := row.Scan(&id); err != nil {
		return 300, err
	}
	return 200, nil
}
func (r *AuthPostgres) GetUser(username, password string) (grod.User, error) {
	var user grod.User
	query := fmt.Sprintf("SELECT id  FROM %s WHERE username=$1 AND password_hash=$2", UserTable)

	err := r.db.Get(&user, query, username, password)
	return user, err
}
func (r *AuthPostgres) UpdateSubTime() error {
	query := fmt.Sprintf("UPDATE %s SET time_of_sub=time_of_sub-1 WHERE time_of_sub <> -1 AND time_of_sub <> 0", UserTable)
	_, err := r.db.Exec(query)
	return err
}
