package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/xxxle0/tribee-admin-backend/models"
)

var Db *sqlx.DB

func findAll() ([]models.User, error) {
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

func findById(id string) (models.User, error) {
	user := models.User{}
	err := Db.Select(&user, "SELECT * FROM Users WHERE id = ?", id)
	if err != nil {
		return user, errors.Wrap(err, "unable to get hour from db")
	}
	return user, nil
}
