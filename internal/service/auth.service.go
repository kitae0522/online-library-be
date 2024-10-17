package service

import (
	"github.com/kitae0522/online-library-be/internal/repository"
	"github.com/kitae0522/online-library-be/pkg/crypt"
	"github.com/kitae0522/online-library-be/pkg/domain"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: repo}
}

func (s *AuthService) Register(req *domain.AuthRegisterReq) error {
	// 1. Compare Password, PasswordConfirm
	// -> boolean (true: is same password, false: is not same password)
	if req.Password != req.PasswordConfirm {
		return domain.ErrIncorrectConfirmPassword
	}

	// 2. Create User
	user, err := s.authRepo.CreateUser(req)
	if err != nil {
		return err
	}

	// 3. Create UserPassword
	if err := s.authRepo.CreateUserPassword(user, req.Password); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(req *domain.AuthLoginReq) (string, error) {
	// 1. Get UserInfo By Email
	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	// 2. Get PasswordInfo
	passwordInfo, err := s.authRepo.GetUserPassword(user)
	if err != nil {
		return "", err
	}

	// 3. Compare Password from DB, InputPassword
	// -> boolean (true: is same password, false: is not same password)
	if !crypt.VerifyPassword(passwordInfo.Password, req.Password, passwordInfo.Salt) {
		return "", domain.ErrWrongPassword
	}

	// 4. Generate JWT Token
	token, err := crypt.NewToken(string(user.Role), user.UserUUID, []byte("tempSecret"))
	return token, err
}

func (s *AuthService) PasswordReset(req *domain.AuthPasswordResetReq) error {
	// 1. Compare NewPassword, NewPasswordConfirm
	// -> boolean (true: is same password, false: is not same password)
	if req.NewPassword != req.NewPasswordConfirm {
		return domain.ErrIncorrectConfirmPassword
	}

	// 2. Get UserInfo
	user, err := s.authRepo.GetUserByUUID(req.UserUUID)
	if err != nil {
		return err
	}

	// 3. Get UserPassword
	passwordInfo, err := s.authRepo.GetUserPassword(user)
	if err != nil {
		return err
	}

	// 4. Compare PasswordInfo from DB, OldPassword
	// -> boolean (true: is same password, false: is not same password)
	if !crypt.VerifyPassword(passwordInfo.Password, req.OldPassword, passwordInfo.Salt) {
		return domain.ErrWrongPassword
	}

	// 5. Update UserPassword
	if err := s.authRepo.UpdateUserPassword(user, req.NewPassword); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Withdraw(req *domain.AuthWithdrawReq) error {
	// 1. Get UserInfo By UUID
	user, err := s.authRepo.GetUserByUUID(req.UserUUID)
	if err != nil {
		return err
	}

	// 2. Delete User
	if ok, err := s.authRepo.DeleteUser(user); err != nil {
		return err
	} else if !ok {
		return domain.ErrUnableToDeleteUser
	}
	return nil
}
