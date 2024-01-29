package user_usecase

type UserUseCases struct {
	RegisterUser   RegisterUserUseCase
	UpdateProfile  UpdateUserImageUsecase
	LoginUser      LoginUserUsecase
	SendAuthEmail  *SendAuthEmailUseCase
	VerifyPasscode VerifyPasscodeUseCase
	LogoutUser     LogoutUserUseCase
}

func NewUserCases(registerUserUsecase RegisterUserUseCase, updtProfUsecs UpdateUserImageUsecase, loginUser LoginUserUsecase, sendEmail *SendAuthEmailUseCase, verifyPass VerifyPasscodeUseCase, logout LogoutUserUseCase) *UserUseCases {

	return &UserUseCases{
		RegisterUser:   registerUserUsecase,
		UpdateProfile:  updtProfUsecs,
		LoginUser:      loginUser,
		SendAuthEmail:  sendEmail,
		VerifyPasscode: verifyPass,
		LogoutUser:     logout,
	}
}
