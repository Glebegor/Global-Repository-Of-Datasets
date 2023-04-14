package service

import "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"

type DatasetItemsService struct {
	repo repository.DatasetItems
}

func NewDatasetItemsService(repo repository.DatasetItems) *DatasetItemsService {
	return &DatasetItemsService{repo: repo}
}
func (r *DatasetItemsService) GetAll(userId int) (int, error) {
	return 0, nil
}
