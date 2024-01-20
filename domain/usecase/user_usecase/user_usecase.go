package user_usecase

type UserUseCases struct {
	RegisterUser  RegisterUserUseCase
	UpdateProfile UpdateUserImageUsecase
	LoginUser     LoginUserUsecase
}

func NewUserCases(registerUserUsecase RegisterUserUseCase, updtProfUsecs UpdateUserImageUsecase, loginUser LoginUserUsecase) *UserUseCases {

	return &UserUseCases{
		RegisterUser:  registerUserUsecase,
		UpdateProfile: updtProfUsecs,
		LoginUser:     loginUser,
	}
}
