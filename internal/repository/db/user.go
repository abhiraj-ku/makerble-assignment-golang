package postresdb

import (
	"database/sql"
	"fmt"
	"log"

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

func (r *UserRepo) FindByUsername(name string) (*model.User, error) {
	user := &model.User{}
	stmt, err := r.db.Prepare(`select id,name,password,role ,created_at from users where name = $1`)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by name %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&user.ID, &user.Name, &user.Password, &user.Role, &user.CreatedAt)
	log.Print(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("failed to populate the data in findByUser %w", err)
		}
	}
	return user, nil

}

func (r *UserRepo) CreateUser(user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (name, password, role,created_at)
		VALUES ($1, $2, $3,$4)
		RETURNING id, name, role, created_at;
	`

	var created model.User
	// user.CreatedAt = time.Now()

	err := r.db.QueryRow(query, user.Name, user.Password, user.Role, user.CreatedAt).Scan(&created.ID, &created.Name, &created.Role, &created.CreatedAt)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("coudn't create error %w", err)
	}
	log.Println(created)
	return &created, nil
}
