package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/krinya32/todoApp"
	"github.com/krinya32/todoApp/pkg/repository"
)

const salt = "hjgrhjqw123[sgrgop4m"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todoApp.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
