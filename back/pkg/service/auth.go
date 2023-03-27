package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (r *AuthService) CreateUser(user grod.User) (int, error) {
	user.Password = r.genPasswordHash(user.Password)
	return r.repo.CreateUser(user)
}

func (r *AuthService) genPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("salt"))))
}
