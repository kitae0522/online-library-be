package repository

import (
	"github.com/kitae0522/online-library-be/internal/model"
)

type UserRepository struct {
	client *model.PrismaClient
}

func NewUserRepository(prismaClient *model.PrismaClient) *UserRepository {
	return &UserRepository{client: prismaClient}
}
