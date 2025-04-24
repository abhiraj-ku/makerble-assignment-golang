package repository

import "github.com/abhiraj-ku/health_app/internal/model"

type UserRepository interface {
	FindByUsername(name string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}
