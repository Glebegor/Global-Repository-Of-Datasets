package repository

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/jmoiron/sqlx"
)

const (
	UserTable          = "users"
	UsersDatasetsTable = "users_datasets"
	DatasetsItemsTable = "datasets_items"
	DatasetItemTable   = "datasetItem"
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
type Datasets interface {
	Create(userId int, dataset grod.Dataset) error
	GetAll(userId int) ([]grod.Dataset, error)
	GetById(userId, datasetId int) (grod.Dataset, error)
	GetRandom(userId, datasetId int) (grod.DatasetItem, error)
	Delete(userId, datasetId int) error
	Update(userId, datasetId int, input grod.UpdateDataset) error
}
type DatasetItems interface {
	Create(userId int, datasetId int, data grod.DatasetItem) error
	GetAll(usersId int, datasetId int) ([]grod.DatasetItem, error)
}
type Repository struct {
	Authorization
	Subscribe
	Datasets
	DatasetItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Subscribe:     NewSubscribePostgres(db),
		Datasets:      NewDatasetsPostgres(db),
		DatasetItems:  NewDatasetItemsPostgres(db),
	}
}
