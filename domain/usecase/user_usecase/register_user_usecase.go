package user_usecase

import (
	"fmt"

	"github.com/proGabby/4genz/domain/entity"
	"github.com/proGabby/4genz/domain/repository/user_repo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	userRepo user_repo.UserRepository
}

func (u *RegisterUserUseCase) Execute(name, email, password string) (*entity.User, error) {

	isExist, err := u.userRepo.CheckIfUserExists(email)

	if err != nil {
		return nil, err
	}

	if isExist {
		return nil, fmt.Errorf("User already exists")
	}

	hashedPassword, er := u.hashedPassword(password)
	if er != nil {
		return nil, er
	}

	return u.userRepo.RegisterUser(name, email, hashedPassword)
}

func (u *RegisterUserUseCase) hashedPassword(password string) ([]byte, error) {
	hashedPassword, er := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if er != nil {
		return nil, er
	}

	return hashedPassword, nil
}

func NewRegisterUserUseCase(userRepo user_repo.UserRepository) *RegisterUserUseCase {

	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}
