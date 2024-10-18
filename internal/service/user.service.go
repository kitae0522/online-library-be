package service

import (
	"github.com/kitae0522/online-library-be/internal/repository"
	"github.com/kitae0522/online-library-be/pkg/domain"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUserProfile(req *domain.UserCreateUserProfileReq) error {
	// 1. Create UserProfile
	if _, err := s.userRepo.CreateUserProfile(req); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserUUIDByTag(userTag string) (*domain.UserGetUserProfileEntity, error) {
	// 1. Get UserUUID by userTag
	user, err := s.userRepo.GetUserUUIDByTag(userTag)
	if err != nil {
		return nil, err
	}

	// 2. Get UserProfile by userUUID
	profile, err := s.userRepo.GetUserProfileByUUID(user.UserUUID)
	if err != nil {
		return nil, err
	}

	// 3. Initialize ProfilePic, Bio with default value
	profilePic, bio := "", ""
	if p, ok := profile.ProfilePic(); ok {
		profilePic = p
	}
	if b, ok := profile.Bio(); ok {
		bio = b
	}

	// 4. Create UserGetUserProfileEntity
	profileEntity := &domain.UserGetUserProfileEntity{
		UserUUID:   user.UserUUID,
		UserTag:    user.UserTag,
		Email:      user.Email,
		Name:       user.Name,
		ProfilePic: profilePic,
		Bio:        bio,
		Role:       string(user.Role),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return profileEntity, nil
}
