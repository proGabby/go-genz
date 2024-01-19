package user_repo_impl

import (
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/domain/entity"
)

type UserRepositoryImpl struct {
	psql postgressDatasource.PostgresDBStore
}

// UpdateProfileImage implements user_repo.UserRepository.
func (*UserRepositoryImpl) UpdateProfileImage(userId int, profileImageUrl string) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserRepoImpl(psql postgressDatasource.PostgresDBStore) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		psql: psql,
	}
}

func (userRepoImpl *UserRepositoryImpl) RegisterUser(name, email string, hashedPassword []byte) (*entity.User, error) {

	return userRepoImpl.psql.RegisterUser(name, email, "https://genzstorage.blob.core.windows.net/genzblob/user_person_human_avatar-128.webp", hashedPassword)
}

func (userRepoImpl *UserRepositoryImpl) UpdateUser(userID int, username, password, profileImageUrl string) (*entity.User, error) {
	panic("not implemted")
}

func (userRepoImpl *UserRepositoryImpl) DeleteUser(userID int) error {
	panic("not implemted")
}

func (userRepoImpl *UserRepositoryImpl) GetUserByID(userID int) (*entity.User, error) {
	panic("not implemted")
}

func (userRepoImpl *UserRepositoryImpl) VerifyUserCredentials(username, password string) (*entity.User, error) {
	panic("not implemted")
}
