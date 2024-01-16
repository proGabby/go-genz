package user_usecase

import "github.com/proGabby/4genz/domain/repository/user_repo"

type UserUseCase struct {
	UserRepository user_repo.UserRepository; 
}