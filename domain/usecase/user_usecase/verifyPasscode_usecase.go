package user_usecase

import (
	"fmt"

	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/repository/user_repo"
)

type VerifyPasscodeUseCase struct {
	userRepo user_repo.UserRepository
}

func NewVerifyPasscodeUseCase(userRepo user_repo.UserRepository) *VerifyPasscodeUseCase {

	return &VerifyPasscodeUseCase{
		userRepo: userRepo,
	}
}

func (c *VerifyPasscodeUseCase) Execute(userId int, passcode string) (*dto.UserResponse, error) {

	user, err := c.userRepo.FetchUserPasscode(userId)

	if err != nil {
		fmt.Printf("fetching user passcode err: %v", err)
		return nil, err
	}

	if user == nil {
		fmt.Printf("user not found")
		return nil, fmt.Errorf("user not found")
	}

	if user.IsVerified {
		fmt.Printf("user already verified")
		return nil, fmt.Errorf("user already verified")
	}

	if user.EmailOtp != passcode {
		return nil, fmt.Errorf("passcode does not match")
	}

	userRes, err := c.userRepo.UpdateUserEmailVerificationStatus(userId, true)

	return userRes, err
}
