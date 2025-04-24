package postresdb

import (
	"database/sql"

	"github.com/abhiraj-ku/health_app/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query, err := r.db.Prepare(`select id,username,role ,created_at from users where username = $1`)
	if err != nil {
		return &model.User{}, err
	}
	defer query.Close()

	err = r.db.QueryRow(username).Scan(&user.ID, &user.Name, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &model.User{}, err
		}
	}
	return user, nil

}
