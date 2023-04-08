package service

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
)

type Authorization interface {
	CreateUser(user grod.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accesToken string) (int, string, error)
}

type Subscribe interface {
	BuyCommon(user_id int) (int, error)
	UnSubCommon(user_id int) (int, error)
	BuyPremium(user_id int) (int, error)
	UnSubPremium(user_id int) (int, error)
}
type Datasets interface {
	Create(userId int, input grod.Dataset) error
	GetAll(userId int) ([]grod.Dataset, error)
	GetById(userId, datasetId int) (grod.Dataset, error)
	Delete(userId, datasetId int) error
}
type Service struct {
	Authorization
	Subscribe
	Datasets
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Subscribe:     NewSubscribeService(repos.Subscribe),
		Datasets:      NewDatasetsService(repos.Datasets),
	}
}
