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
func (r *DatasetsService) GetById(userId, datasetId int) (grod.Dataset, error) {
	data, err := r.repo.GetById(userId, datasetId)
	return data, err
}
func (r *DatasetsService) Delete(userId, datasetId int) error {
	err := r.repo.Delete(userId, datasetId)
	return err
}
