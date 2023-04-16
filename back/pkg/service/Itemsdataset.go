package service

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
)

type DatasetItemsService struct {
	repo repository.DatasetItems
}

func NewDatasetItemsService(repo repository.DatasetItems) *DatasetItemsService {
	return &DatasetItemsService{repo: repo}
}
func (r *DatasetItemsService) Create(userId int, datasetId int, data grod.DatasetItem) error {
	err := r.repo.Create(userId, datasetId, data)
	return err
}
func (r *DatasetItemsService) GetAll(userId int, datasetId int) ([]grod.DatasetItem, error) {
	data, err := r.repo.GetAll(userId, datasetId)
	return data, err
}
func (r *DatasetItemsService) ItemsGet(userId, datasetId, itemId int) (grod.DatasetItem, error) {
	data, err := r.repo.ItemsGet(userId, datasetId, itemId)
	return data, err
}
func (r *DatasetItemsService) ItemsDelete(userId, datasetId, itemId int) error {
	err := r.repo.ItemsDelete(userId, datasetId, itemId)
	return err
}
func (r *DatasetItemsService) ItemsUpdate(userId int, datasetId int, itemId int, input grod.UpdateDatasetItem) error {
	err := r.repo.ItemsUpdate(userId, datasetId, itemId, input)
	return err
}
