package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/xxxle0/tribee-admin-backend/models"
)

var Db *sqlx.DB

type UserRepository struct{}

func (repository *UserRepository) FindAll() ([]models.User, error) {
	result := []models.User{}
	rows, err := Db.Query(`SELECT * FROM Users`)
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

func (repository *UserRepository) FindById(id string) (models.User, error) {
	user := models.User{}
	err := Db.Select(&user, "SELECT * FROM Users WHERE id = ?", id)
	if err != nil {
		return user, errors.Wrap(err, "Get user error")
	}
	return user, nil
}

func (repository *UserRepository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := Db.Select(&user, "SELECT * FROM Users WHERE email = ?", email)
	if err != nil {
		return user, errors.Wrap(err, "Get user error")
	}
	return user, nil
}
