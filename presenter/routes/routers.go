package userRoutes

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"

	infrastruture "github.com/proGabby/4genz/data/cloud_infrastruture"
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	cloud_repo_impl "github.com/proGabby/4genz/data/repo_impl/cloub_repo_impl"
	"github.com/proGabby/4genz/data/repo_impl/user_repo_impl.go"
	"github.com/proGabby/4genz/domain/usecase/user_usecase"
	"github.com/proGabby/4genz/presenter/controllers"
	"github.com/proGabby/4genz/presenter/middlewares"
)

func SetUpUserRoutes(r *mux.Router, db *sql.DB) {

	psql := postgressDatasource.NewPostgresDBStore(db)
	// azCloud, err := infrastruture.NewAzureCloudInfrasture()
	// if err != nil {
	// 	return
	// }

	awsStorage, err := infrastruture.NewS3Handler()
	if err != nil {
		fmt.Printf("AWS storage error: %v", err)
		return
	}

	userRepoImpl := user_repo_impl.NewUserRepoImpl(*psql)
	// cloudStrImpl := cloud_repo_impl.NewCloudReposityImpl(*azCloud)
	cloudStrImpl := cloud_repo_impl.NewAWSCloudReposityImpl(*awsStorage)
	emailSendImpl, err := infrastruture.NewGomailSender("smtp.gmail.com")
	var sendEmailUsecase *user_usecase.SendAuthEmailUseCase
	if err != nil {
		fmt.Printf("Email service error: %v", err)
	}

	registerUserUsecase := user_usecase.NewRegisterUserUseCase(userRepoImpl)
	updateUserImage := user_usecase.NewUpdateUserImageUsecase(cloudStrImpl, userRepoImpl)
	loginUserUsecase := user_usecase.NewLoginUserUsecase(userRepoImpl)
	sendEmailUsecase = user_usecase.NewSendAuthEmailUseCase(userRepoImpl, emailSendImpl)
	verifyPasscodeUsecase := user_usecase.NewVerifyPasscodeUseCase(userRepoImpl)
	logoutUser := user_usecase.NewLogoutUserUseCase(userRepoImpl)

	userUsecases := user_usecase.NewUserCases(*registerUserUsecase, *updateUserImage, *loginUserUsecase, sendEmailUsecase, *verifyPasscodeUsecase, *logoutUser)
	userController := controllers.NewUserController(*userUsecases)

	authorizer := middlewares.NewAuthMiddleware(*userRepoImpl)

	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/update/profile-img", authorizer.Authenticate(userController.UpadateUserImage)).Methods("POST")
	r.HandleFunc("/login", userController.UserLogin).Methods("POST")
	r.HandleFunc("/send-auth-email", authorizer.Authenticate(userController.SendAuthEmail)).Methods("POST")
	r.HandleFunc("/verify-passcode", authorizer.Authenticate(userController.VerifyPasscode)).Methods("POST")
	r.HandleFunc("/logout", authorizer.Authenticate(userController.LogoutUser)).Methods("DELETE")
}
