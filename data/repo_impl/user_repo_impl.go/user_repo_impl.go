package user_repo_impl

import (
	"fmt"

	postgressDatasource "github.com/proGabby/4genz/data/datasource"
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

func (userRepoImpl *UserRepositoryImpl) RegisterUser(name, email, profileImageUrl string, hashedPassword []byte) (*entity.User, error) {
	fmt.Println("recorded")
	return userRepoImpl.psql.RegisterUser(name, email, profileImageUrl, hashedPassword)
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
