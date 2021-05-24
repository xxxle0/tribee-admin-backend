package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/xxxle0/tribee-admin-backend/models"
)

func get(conn *sqlx.DB) (models.User, error) {
	user := models.User{}
	err := conn.Select(&user, `SELECT * FROM Users`)
	if err != nil {
		return user, errors.Wrap(err, "unable to get hour from db")
	}
	return user, nil
}
