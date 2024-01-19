package userRoutes

import (
	"database/sql"

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
	CloudStrImpl := cloud_repo_impl.NewCloudReposityImpl(*azCloud)

	registerUserUsecase := user_usecase.NewRegisterUserUseCase(userRepoImpl)
	updateUserImage := user_usecase.NewUpdateUserImageUsecase(CloudStrImpl, userRepoImpl)

	userUsecases := user_usecase.NewUserCases(*registerUserUsecase, *updateUserImage)
	userController := controllers.NewUserController(*userUsecases)

	authorizer := middlewares.NewAuthMiddleware(*userRepoImpl)

	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/update/profile-img", authorizer.Authenticate(userController.UpadateUserImage)).Methods("POST")
}
