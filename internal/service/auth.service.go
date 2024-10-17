package service

import (
	"github.com/kitae0522/online-library-be/internal/dto"
	"github.com/kitae0522/online-library-be/internal/repository"
	"github.com/kitae0522/online-library-be/pkg/crypt"
	"github.com/kitae0522/online-library-be/pkg/utils"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: repo}
}

func (s *AuthService) Register(req *dto.AuthRegisterReq) error {
	user, err := s.authRepo.CreateUser(req)
	if err != nil {
		return err
	}
	if err := s.authRepo.CreateUserPassword(user, req.Password); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Login(req *dto.AuthLoginReq) (string, error) {
	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	passwordInfo, err := s.authRepo.GetUserPassword(user)
	if err != nil {
		return "", err
	}

	if !crypt.VerifyPassword(passwordInfo.Password, req.Password, crypt.EncodeBase64(user.UserUUID)) {
		return "", utils.ErrWrongPassword
	}

	token, err := crypt.NewToken(string(user.Role), user.UserUUID, []byte("tempSecret"))
	return token, err
}
