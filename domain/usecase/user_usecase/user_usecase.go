package user_usecase

type UserUseCases struct {
	RegisterUser  RegisterUserUseCase
	UpdateProfile UpdateUserImageUsecase
}

func NewUserCases(registerUserUsecase RegisterUserUseCase, updtProfUsecs UpdateUserImageUsecase) *UserUseCases {

	return &UserUseCases{
		RegisterUser:  registerUserUsecase,
		UpdateProfile: updtProfUsecs,
	}
}
