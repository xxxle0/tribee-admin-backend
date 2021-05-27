package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/xxxle0/tribee-admin-backend/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

type UserRepositoryI interface {
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

func UserRepositoryInit(db *sqlx.DB) UserRepositoryI {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	result := []models.User{}
	rows, err := r.DB.Query(`SELECT * FROM Users`)
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Email, &user.Password)
		if err != nil {
			log.Fatal("scan error", err)
		}
		result = append(result, user)
	}
	return result, nil
}

func (r *UserRepository) FindById(id string) (models.User, error) {
	user := models.User{}
	err := r.DB.Select(&user, "SELECT * FROM Users WHERE id = ?", id)
	if err != nil {
		return user, errors.Wrap(err, "Get user error")
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.DB.Select(&user, "SELECT * FROM Users WHERE email = ?", email)
	if err != nil {
		return user, errors.Wrap(err, "Get user error")
	}
	return user, nil
}
