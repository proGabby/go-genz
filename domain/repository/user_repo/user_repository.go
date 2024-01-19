package user_repo

import "github.com/proGabby/4genz/domain/entity"

type UserRepository interface {
	RegisterUser(name, email string, password []byte) (*entity.User, error)
	// UpdateUser(userID int, username, password, profileImageUrl string) (*entity.User, error)
	UpdateProfileImage(userId int, profileImageUrl string) (*entity.User, error)
	DeleteUser(userID int) error
	GetUserByID(userID int) (*entity.User, error)
	VerifyUserCredentials(username, password string) (*entity.User, error)
}
