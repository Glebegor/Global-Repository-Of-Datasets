package repository

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/jmoiron/sqlx"
)

const (
	UserTable          = "users"
	UsersDatasetsTable = "users_datasets"
	DatasetsItemsTable = "datasets_items"
	DatasetItemTable   = "dataset_item"
	DatasetsTable      = "datasets"
)

type Authorization interface {
	CreateUser(user grod.User) (int, error)
	GetUser(username, password string) (grod.User, error)
}
type Subscribe interface {
	BuyCommon(user_id int) (int, error)
	UnSubCommon(user_id int) (int, error)
	BuyPremium(user_id int) (int, error)
	UnSubPremium(user_id int) (int, error)
}
type Repository struct {
	Authorization
	Subscribe
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Subscribe:     NewSubscribePostgres(db),
	}
}
