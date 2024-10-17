package repository

import (
	"context"

	"github.com/kitae0522/online-library-be/internal/dto"
	"github.com/kitae0522/online-library-be/internal/model"
	"github.com/kitae0522/online-library-be/pkg/crypt"
)

type AuthRepository struct {
	client *model.PrismaClient
}

func NewAuthRepository(prismaClient *model.PrismaClient) *AuthRepository {
	return &AuthRepository{client: prismaClient}
}

func (r *AuthRepository) CreateUser(req *dto.AuthRegisterReq) (*model.UsersModel, error) {
	user, err := r.client.Users.CreateOne(
		model.Users.UserTag.Set(req.UserTag),
		model.Users.Email.Set(req.Email),
		model.Users.Role.Set(model.UserRolesUser),
		model.Users.Name.Set(req.Name),
	).Exec(context.Background())
	return user, err
}

func (r *AuthRepository) CreateUserPassword(user *model.UsersModel, password string) error {
	salt := crypt.EncodeBase64(user.UserUUID)
	hashedPassword := crypt.NewSHA256(password, salt)

	_, err := r.client.UserPassword.CreateOne(
		model.UserPassword.Password.Set(hashedPassword),
		model.UserPassword.Salt.Set(salt),
		model.UserPassword.User.Link(
			model.Users.UserUUID.Equals(user.UserUUID),
		),
	).Exec(context.Background())
	return err
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.UsersModel, error) {
	user, err := r.client.Users.FindUnique(
		model.Users.Email.Equals(email),
	).Exec(context.Background())
	return user, err
}

func (r *AuthRepository) GetUserByUUID(uuid string) (*model.UsersModel, error) {
	user, err := r.client.Users.FindUnique(
		model.Users.UserUUID.Equals(uuid),
	).Exec(context.Background())
	return user, err
}

func (r *AuthRepository) GetUserPassword(user *model.UsersModel) (*model.UserPasswordModel, error) {
	passwordInfo, err := r.client.UserPassword.FindUnique(
		model.UserPassword.UserUUID.Equals(user.UserUUID),
	).Exec(context.Background())
	return passwordInfo, err
}
