package repository

import (
	"context"

	"github.com/kitae0522/online-library-be/internal/model"
	"github.com/kitae0522/online-library-be/pkg/domain"
)

type UserRepository struct {
	client *model.PrismaClient
}

func NewUserRepository(prismaClient *model.PrismaClient) *UserRepository {
	return &UserRepository{client: prismaClient}
}

func (r *UserRepository) CreateUserProfile(req *domain.UserCreateUserProfileReq) (*model.UserProfileModel, error) {
	profile, err := r.client.UserProfile.CreateOne(
		model.UserProfile.User.Link(
			model.Users.UserUUID.Equals(req.UserUUID),
		),
		model.UserProfile.ProfilePic.SetIfPresent(req.ProfilePic),
		model.UserProfile.Bio.SetIfPresent(req.Bio),
	).Exec(context.Background())
	return profile, err
}

func (r *UserRepository) GetUserUUIDByTag(userTag string) (*model.UsersModel, error) {
	user, err := r.client.Users.FindUnique(
		model.Users.UserTag.Equals(userTag),
	).Exec(context.Background())
	return user, err
}

func (r *UserRepository) GetUserProfileByUUID(uuid string) (*model.UserProfileModel, error) {
	profile, err := r.client.UserProfile.FindUnique(
		model.UserProfile.UserUUID.Equals(uuid),
	).Exec(context.Background())
	return profile, err
}
