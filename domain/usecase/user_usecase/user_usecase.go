package user_usecase

type UserUseCases struct {
	RegisterUser  RegisterUserUseCase
	UpdateProfile UpdateUserImageUsecase
	LoginUser     LoginUserUsecase
	SendAuthEmail *SendAuthEmailUseCase
}

func NewUserCases(registerUserUsecase RegisterUserUseCase, updtProfUsecs UpdateUserImageUsecase, loginUser LoginUserUsecase, sendEmail *SendAuthEmailUseCase) *UserUseCases {

	return &UserUseCases{
		RegisterUser:  registerUserUsecase,
		UpdateProfile: updtProfUsecs,
		LoginUser:     loginUser,
		SendAuthEmail: sendEmail,
	}
}
