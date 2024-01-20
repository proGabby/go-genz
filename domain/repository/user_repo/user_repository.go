package user_repo

import (
	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/entity"
)

type UserRepository interface {
	RegisterUser(name, email string, password []byte) (*entity.User, error)
	UpdateUser(userID int, username, password, profileImageUrl string) (*entity.User, error)
	UpdateProfileImage(userId int, profileImageUrl string) (*dto.UserResponse, error)
	DeleteUser(userID int) error
	GetUserByID(userID int) (*entity.User, error)
	VerifyUserCredentials(email string) (*entity.User, error)
}
