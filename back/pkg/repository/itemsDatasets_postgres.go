package repository

import (
	"fmt"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/jmoiron/sqlx"
)

type DatasetItemsPostgres struct {
	db *sqlx.DB
}

func NewDatasetItemsPostgres(db *sqlx.DB) *DatasetItemsPostgres {
	return &DatasetItemsPostgres{db: db}
}
func (r *DatasetItemsPostgres) Create(userId int, datasetId int, data grod.DatasetItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createDatasetQuery := fmt.Sprintf("INSERT INTO %s (datainfo, solution) VALUES ($1, $2) RETURNING id", DatasetItemTable)
	row := tx.QueryRow(createDatasetQuery, data.Datainfo, data.Solution)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	createDatasetUserQuery := fmt.Sprintf("INSERT INTO %s (id_dataset, id_item) VALUES ($1, $2)", DatasetsItemsTable)
	_, err = tx.Exec(createDatasetUserQuery, datasetId, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *DatasetItemsPostgres) GetAll(usersId int, datasetId int) ([]grod.DatasetItem, error) {
	var data []grod.DatasetItem

	query := fmt.Sprintf("SELECT tl.id, tl.solution, tl.datainfo FROM %s tl INNER JOIN %s ul on tl.id=ul.id_item WHERE id_dataset=$1", DatasetItemTable, DatasetsItemsTable)
	err := r.db.Select(&data, query, datasetId)
	return data, err
}
