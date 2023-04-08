package service

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
)

type DatasetsService struct {
	repo repository.Datasets
}

func NewDatasetsService(repo repository.Datasets) *DatasetsService {
	return &DatasetsService{repo: repo}
}

func (r *DatasetsService) Create(userId int, input grod.Dataset) error {
	err := r.repo.Create(userId, input)
	return err
}
func (r *DatasetsService) GetAll(userId int) ([]grod.Dataset, error) {
	data, err := r.repo.GetAll(userId)
	return data, err
}