package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name               string `db:"name"`
	Email              string `db:"email"`
	PhoneNumber        string `db:"phone_number"`
	Password           string `db:"password"`
	Avatar             string `db:"avatar"`
	Birthday           string `db:"birthday"`
	RecognitionReceive string `db:"recognition_receive`
	RecognitionSend    string `db:"recognition_send`
}

func (u *User) HashingPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CompareHashedPassword(hashedPassword string, originalPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(originalPassword))
	return err == nil
}
