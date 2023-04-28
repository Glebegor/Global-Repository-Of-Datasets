package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SubscribePostgres struct {
	db *sqlx.DB
}

func NewSubscribePostgres(db *sqlx.DB) *SubscribePostgres {
	return &SubscribePostgres{db: db}
}

func (r *SubscribePostgres) BuyCommon(user_id int) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET subscribe='common', time_of_sub=$1 WHERE id=$2", UserTable)
	_, err := r.db.Exec(query, uint(30*24), user_id)

	return 200, err
}
func (r *SubscribePostgres) UnSubCommon(user_id int) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET subscribe='free', time_of_sub=-1 WHERE id=$1", UserTable)
	_, err := r.db.Exec(query, user_id)

	return 200, err
}
func (r *SubscribePostgres) BuyPremium(user_id int) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET subscribe='premium', time_of_sub=$1 WHERE id=$2", UserTable)
	_, err := r.db.Exec(query, uint(30*24), user_id)

	return 200, err
}
func (r *SubscribePostgres) UnSubPremium(user_id int) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET subscribe='free', time_of_sub=-1 WHERE id=$1", UserTable)
	_, err := r.db.Exec(query, user_id)

	return 200, err
}
