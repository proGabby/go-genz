package user_usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/repository/user_repo"
)

type LoginUserUsecase struct {
	userRepo user_repo.UserRepository
}

func NewLoginUserUsecase(userRepo user_repo.UserRepository) *LoginUserUsecase {
	return &LoginUserUsecase{
		userRepo: userRepo,
	}
}

func (u *LoginUserUsecase) Execute(email, password string) (map[string]interface{}, error) {
	user, err := u.userRepo.VerifyUserCredentials(email)

	if err != nil {
		return nil, err
	}

	error := u.compareHashedPassword(user.Password, password)

	if error != nil {
		return nil, error
	}

	token, err := u.generateJWTToken(user.Id)

	if error != nil {
		return nil, error
	}

	userRes := dto.UserResponse{
		Id:              user.Id,
		Name:            user.Name,
		Email:           user.Email,
		ProfileImageUrl: user.ProfileImageUrl,
	}

	return map[string]interface{}{
		"user":  userRes,
		"token": token,
	}, nil

}

func (u *LoginUserUsecase) compareHashedPassword(hashedPassword, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Print(err)
		return errors.New("incorrect password")
	}
	return nil
}

func (u *LoginUserUsecase) generateJWTToken(userId int) (string, error) {

	secrtKey, err := getJWTSecretKey()

	if err != nil {
		return "", err
	}

	// Define the expiration time for the token (e.g., 1 hour)
	expirationTime := time.Now().Add(12 * time.Hour)

	// Create the JWT claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   fmt.Sprintf("%d", userId),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte(secrtKey) // Replace with a secure secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getJWTSecretKey() (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY not set")
	}
	return secretKey, nil
}
