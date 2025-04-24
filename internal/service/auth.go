package service

import (
	"errors"

	"github.com/abhiraj-ku/health_app/internal/model"
	"github.com/abhiraj-ku/health_app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Authenticate(username string, password string) (*model.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username, this username doesn't exist ")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password does not match, Try again")
	}

	return user, nil
}
