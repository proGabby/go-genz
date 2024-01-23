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
	azCloud, err := infrastruture.NewAzureCloudInfrasture()
	if err != nil {
		return
	}

	userRepoImpl := user_repo_impl.NewUserRepoImpl(*psql)
	cloudStrImpl := cloud_repo_impl.NewCloudReposityImpl(*azCloud)
	emailSendImpl, err := infrastruture.NewGomailSender("smtp.gmail.com")
	var sendEmailUsecase *user_usecase.SendAuthEmailUseCase
	if err != nil {
		fmt.Printf("Email service error: %v", err)
	}
	sendEmailUsecase = user_usecase.NewSendAuthEmailUseCase(userRepoImpl, emailSendImpl)

	registerUserUsecase := user_usecase.NewRegisterUserUseCase(userRepoImpl)
	updateUserImage := user_usecase.NewUpdateUserImageUsecase(cloudStrImpl, userRepoImpl)
	loginUserUsecase := user_usecase.NewLoginUserUsecase(userRepoImpl)

	userUsecases := user_usecase.NewUserCases(*registerUserUsecase, *updateUserImage, *loginUserUsecase, sendEmailUsecase)
	userController := controllers.NewUserController(*userUsecases)

	authorizer := middlewares.NewAuthMiddleware(*userRepoImpl)

	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/update/profile-img", authorizer.Authenticate(userController.UpadateUserImage)).Methods("POST")
	r.HandleFunc("/login", userController.UserLogin).Methods("POST")
	r.HandleFunc("/send-auth-email", authorizer.Authenticate(userController.SendAuthEmail)).Methods("POST")
}
