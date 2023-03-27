package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId       int    `json:"user_id"`
	UserUsername string `json:"user_username"`
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

func (r *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := r.repo.GetUser(username, r.genPasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		username,
	})
	return token.SignedString([]byte(os.Getenv("secretKey")))
}
