package service

import "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
