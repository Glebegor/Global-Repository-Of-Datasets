package service

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
)

type AuthDatasets struct {
	repo repository.Datasets
}

func NewDatasetsService(repo repository.Datasets) *AuthDatasets {
	return &AuthDatasets{repo: repo}
}

func (r *AuthDatasets) Create(userId int, input grod.Dataset) error {
	err := r.repo.Create(userId, input)
	return err
}
