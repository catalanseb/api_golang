package repositories

import (
	"api-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetAllUsers() ([]models.Users, error) {
	users := []models.Users{}
	err := u.db.Select(&users, "SELECT id, first_name, last_name, email, created_on FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}

	return users, nil
}
