package repository

import (
	"fmt"
	"strings"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/jmoiron/sqlx"
)

type DatasetsPostgres struct {
	db *sqlx.DB
}

func NewDatasetsPostgres(db *sqlx.DB) *DatasetsPostgres {
	return &DatasetsPostgres{db: db}
}

func (r *DatasetsPostgres) Create(userId int, dataset grod.Dataset) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createDatasetQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", DatasetsTable)
	row := tx.QueryRow(createDatasetQuery, dataset.Title, dataset.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	createDatasetUserQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_dataset) VALUES ($1, $2)", UsersDatasetsTable)
	_, err = tx.Exec(createDatasetUserQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *DatasetsPostgres) GetAll(userId int) ([]grod.Dataset, error) {
	var datasets []grod.Dataset

	query := fmt.Sprintf("SELECT tl.id, tl.description, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.id_dataset WHERE id_user=$1", DatasetsTable, UsersDatasetsTable)
	err := r.db.Select(&datasets, query, userId)
	return datasets, err
}
func (r *DatasetsPostgres) GetById(userId, datasetId int) (grod.Dataset, error) {
	var datasets grod.Dataset
	query := fmt.Sprintf("SELECT tl.id, tl.description, tl.title FROM %s tl INNER JOIN %s ul on tl.id=ul.id_dataset AND ul.id_user=$1 AND ul.id_dataset=$2", DatasetsTable, UsersDatasetsTable)
	err := r.db.Get(&datasets, query, userId, datasetId)
	return datasets, err
}
func (r *DatasetsPostgres) GetRandom(userId, datasetId int) ([]grod.DatasetItem, error) {
	var data []grod.DatasetItem
	query := fmt.Sprintf("SELECT tl.id, tl.datainfo, tl.solution FROM %s tl INNER JOIN %s ul on ul.id_dataset=$1", DatasetItemTable, DatasetsItemsTable)
	err := r.db.Select(&data, query, datasetId)

	return data, err
}
func (r *DatasetsPostgres) Delete(userId, datasetId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl  USING %s ul WHERE tl.id=ul.id_dataset AND ul.id_user=$1 AND ul.id_dataset=$2", DatasetsTable, UsersDatasetsTable)
	_, err := r.db.Exec(query, userId, datasetId)
	return err
}
func (r *DatasetsPostgres) Update(userId, datasetId int, input grod.UpdateDataset) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.id_dataset AND ul.id_dataset=$%d AND ul.id_user=$%d", DatasetsTable, setQuery, UsersDatasetsTable, argId, argId+1)
	args = append(args, datasetId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
