package repository

import (
	"fmt"
	"strings"

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
func (r *DatasetItemsPostgres) ItemsGet(userId, datasetId, itemId int) (grod.DatasetItem, error) {
	var data grod.DatasetItem

	query := fmt.Sprintf("SELECT tl.id, tl.solution, tl.datainfo FROM %s tl INNER JOIN %s ul on tl.id=ul.id_item AND ul.id_dataset=$1 AND ul.id_item=$2", DatasetItemTable, DatasetsItemsTable)
	err := r.db.Get(&data, query, datasetId, itemId)
	return data, err
}
func (r *DatasetItemsPostgres) ItemsDelete(userId, datasetId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.id_item AND ul.id_dataset=$1 AND ul.id_item=$2", DatasetItemTable, DatasetsItemsTable)
	_, err := r.db.Exec(query, datasetId, itemId)
	return err
}
func (r *DatasetItemsPostgres) ItemsUpdate(userId int, datasetId int, itemId int, input grod.UpdateDatasetItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Datainfo != "" {
		setValues = append(setValues, fmt.Sprintf("datainfo=$%d", argId))
		args = append(args, input.Datainfo)
		argId++
	}
	if input.Solution != "" {
		setValues = append(setValues, fmt.Sprintf("solution=$%d", argId))
		args = append(args, input.Solution)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.id_item AND ul.id_item=$%d AND ul.id_dataset=$%d", DatasetItemTable, setQuery, DatasetsItemsTable, argId, argId+1)
	args = append(args, itemId, datasetId)

	_, err := r.db.Exec(query, args...)
	return err
}
