package services

import (
	"log"

	"github.com/xxxle0/tribee-admin-backend/repositories"
	"github.com/xxxle0/tribee-admin-backend/utils"
)

func AdminSerice() string {
	return "huhu"
}

func SignIn(email string, password string) string {
	user, err := repositories.FindByEmail(email)
	if err != nil {
		log.Fatal("huhuhu", err)
	}
	isPasswordCorrect := utils.CompareHashedPassword(user.Password, password)
	if isPasswordCorrect != true {
		log.Fatal("Password is not correct")
	}
	jwtTokenConfig := &utils.JwtWrapper{
		SecretKey:       "test",
		Issuer:          "test",
		ExpirationHours: 1280,
	}
	generatedToken, err := jwtTokenConfig.GenerateJwtToken(email)
	if err != nil {
		log.Fatal("Password is not correct")
	}
	return generatedToken
}
