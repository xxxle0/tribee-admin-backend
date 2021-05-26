package services

import (
	"log"

	"github.com/xxxle0/tribee-admin-backend/repositories"
	"github.com/xxxle0/tribee-admin-backend/utils"
)

type AdminService struct {
	UserRepository repositories.UserRepository
}

func AdminServiceInit(userRepository repositories.UserRepository) AdminService {
	return AdminService{UserRepository: userRepository}
}
func (s *AdminService) SignIn(email string, password string) string {
	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		log.Fatal("Get user error", err)
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
		log.Fatal("Generate jwt token error", err)
	}
	return generatedToken
}
