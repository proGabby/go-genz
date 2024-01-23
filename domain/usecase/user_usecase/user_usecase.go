package user_usecase

type UserUseCases struct {
	RegisterUser   RegisterUserUseCase
	UpdateProfile  UpdateUserImageUsecase
	LoginUser      LoginUserUsecase
	SendAuthEmail  *SendAuthEmailUseCase
	VerifyPasscode VerifyPasscodeUseCase
}

func NewUserCases(registerUserUsecase RegisterUserUseCase, updtProfUsecs UpdateUserImageUsecase, loginUser LoginUserUsecase, sendEmail *SendAuthEmailUseCase, verifyPass VerifyPasscodeUseCase) *UserUseCases {

	return &UserUseCases{
		RegisterUser:   registerUserUsecase,
		UpdateProfile:  updtProfUsecs,
		LoginUser:      loginUser,
		SendAuthEmail:  sendEmail,
		VerifyPasscode: verifyPass,
	}
}
