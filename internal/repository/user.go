package repository

import "github.com/abhiraj-ku/health_app/internal/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
}
