package user_usecase

import "github.com/proGabby/4genz/domain/repository/user_repo"

type LogoutUserUseCase struct {
	userRepo user_repo.UserRepository
}

func NewLogoutUserUseCase(userRepo user_repo.UserRepository) *LogoutUserUseCase {

	return &LogoutUserUseCase{
		userRepo: userRepo,
	}
}

func (u *LogoutUserUseCase) Execute(userId int) error {

	return u.userRepo.UpdateUserTokenVersion(userId)
}
