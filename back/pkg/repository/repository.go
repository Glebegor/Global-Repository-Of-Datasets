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

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
