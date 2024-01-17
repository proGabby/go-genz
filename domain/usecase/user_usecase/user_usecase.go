package user_usecase

type UserUseCases struct {
	RegisterUser RegisterUserUseCase
}

func NewUserCases(registerUserUsecase RegisterUserUseCase) *UserUseCases {

	return &UserUseCases{
		RegisterUser: registerUserUsecase,
	}
}
