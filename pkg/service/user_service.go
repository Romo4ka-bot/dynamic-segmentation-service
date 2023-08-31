package service

import (
	"crypto/sha1"
	"dynamic-segmentation-service/pkg/model"
	"dynamic-segmentation-service/pkg/repository"
	"fmt"
)

const salt = "some-salt"

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	user.HashPassword = generatePasswordHash(user.HashPassword)
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(id int) (model.User, error) {
	return s.repo.GetUser(id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
