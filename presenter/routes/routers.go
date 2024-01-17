package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/data/repo_impl/user_repo_impl.go"
	"github.com/proGabby/4genz/domain/usecase/user_usecase"
	"github.com/proGabby/4genz/presenter/controllers"
)

func SetUpUserRoutes(r *mux.Router, db *sql.DB) {

	psql := postgressDatasource.NewPostgresDBStore(db)
	userRepoImpl := user_repo_impl.NewUserRepoImpl(*psql)
	registerUserUsecase := user_usecase.NewRegisterUserUseCase(userRepoImpl)

	userUsecases := user_usecase.NewUserCases(*registerUserUsecase)
	userController := controllers.NewUserController(*userUsecases)

	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
}
