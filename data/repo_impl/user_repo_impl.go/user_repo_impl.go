package user_repo_impl

import (
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/entity"
)

type UserRepositoryImpl struct {
	psql postgressDatasource.PostgresDBStore
}

func NewUserRepoImpl(psql postgressDatasource.PostgresDBStore) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		psql: psql,
	}
}

func (userRepoImpl *UserRepositoryImpl) RegisterUser(name, email string, hashedPassword []byte) (*entity.User, error) {

	return userRepoImpl.psql.RegisterUser(name, email, "https://genz-blob.s3.amazonaws.com/img/user_person_human_avatar-128.webp", hashedPassword)
}

func (userRepoImpl *UserRepositoryImpl) UpdateUser(userID int, username, password, profileImageUrl string) (*entity.User, error) {
	panic("not implemted")
}

func (userRepoImpl *UserRepositoryImpl) UpdateProfileImage(userId int, profileImageUrl string) (*dto.UserResponse, error) {
	return userRepoImpl.psql.UpdateUserImage(userId, profileImageUrl)
}

func (userRepoImpl *UserRepositoryImpl) DeleteUser(userID int) error {
	panic("not implemted")
}

func (userRepoImpl *UserRepositoryImpl) GetUserByID(userID int) (*entity.User, error) {
	return userRepoImpl.psql.GetUserByID(userID)
}

func (userRepoImpl *UserRepositoryImpl) VerifyUserCredentials(email string) (*entity.User, error) {
	return userRepoImpl.psql.VerifyUserCredentials(email)
}

func (userRepoImpl *UserRepositoryImpl) UpdateUserEmailPasscode(userId int, emailOtp string) (*dto.UserResponse, error) {
	return userRepoImpl.psql.UpdateUserEmailPasscode(userId, emailOtp)
}

func (userRepoImpl *UserRepositoryImpl) FetchUserPasscode(userId int) (*entity.User, error) {
	return userRepoImpl.psql.FetchPasscode(userId)
}

func (userRepoImpl *UserRepositoryImpl) UpdateUserEmailVerificationStatus(userId int, isVerified bool) (*dto.UserResponse, error) {
	return userRepoImpl.psql.UpdateUserEmailVerificationStatus(userId, isVerified)
}
