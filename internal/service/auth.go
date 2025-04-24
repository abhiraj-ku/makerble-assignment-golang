package service

import (
	"errors"
	"log"

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
	log.Println(user)
	if err != nil {
		log.Printf("❌ invalid username, this username doesn't exist ")

		return nil, errors.New("invalid username, this username doesn't exist ")
	}
	log.Println("🔐 Stored hash from DB:", user.Password)
	// log.Println("🔑 Entered password:", password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("❌ Authentication failed for username '%s': password mismatch", username)

		return nil, errors.New("password does not match, Try again")
	}

	return user, nil
}

func (s *AuthService) Register(user *model.User) (*model.User, error) {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	log.Println(string(hashedPassword))
	user.Password = string(hashedPassword)

	// Save to DB
	return s.UserRepo.CreateUser(user)
}
