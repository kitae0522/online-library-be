package service

import (
	"github.com/kitae0522/online-library-be/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}
