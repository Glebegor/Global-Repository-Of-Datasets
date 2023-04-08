package repository

import (
	"fmt"

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
