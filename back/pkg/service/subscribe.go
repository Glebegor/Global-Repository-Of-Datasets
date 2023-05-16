package service

import (
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
)

type SubscribeService struct {
	repo repository.Subscribe
}

func NewSubscribeService(repo repository.Subscribe) *SubscribeService {
	return &SubscribeService{repo: repo}
}
func (r *SubscribeService) BuyCommon(user_id int) (int, error) {
	status, err := r.repo.BuyCommon(user_id)
	return status, err
}
func (r *SubscribeService) UnSubCommon(user_id int) (int, error) {
	status, err := r.repo.UnSubCommon(user_id)
	return status, err
}

func (r *SubscribeService) BuyPremium(user_id int) (int, error) {
	status, err := r.repo.BuyPremium(user_id)
	return status, err
}
func (r *SubscribeService) UnSubPremium(user_id int) (int, error) {
	status, err := r.repo.UnSubPremium(user_id)
	return status, err
}
