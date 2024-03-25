package routes

import (
	"database/sql"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"

	infrastruture "github.com/proGabby/4genz/data/cloud_infrastruture"
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	cloud_repo_impl "github.com/proGabby/4genz/data/repo_impl/cloub_repo_impl"
	"github.com/proGabby/4genz/data/repo_impl/feed_repo_impl"
	"github.com/proGabby/4genz/data/repo_impl/user_repo_impl.go"
	"github.com/proGabby/4genz/domain/entity"
	"github.com/proGabby/4genz/domain/usecase/feeds_usecase"
	"github.com/proGabby/4genz/domain/usecase/user_usecase"
	"github.com/proGabby/4genz/presenter/controllers"
	"github.com/proGabby/4genz/presenter/middlewares"
)

func SetUpUserRoutes(r *mux.Router, db *sql.DB, socketServ *socketio.Server, feedChan *chan *entity.Feed) {

	psql := postgressDatasource.NewPostgresUserDBStore(db)
	fPsql := postgressDatasource.NewPostgresFeedDBStore(db)
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
	feedRepoImpl := feed_repo_impl.NewFeedRepoImpl(*fPsql, socketServ)
	// cloudStrImpl := cloud_repo_impl.NewCloudReposityImpl(*azCloud)
	cloudStrImpl := cloud_repo_impl.NewAWSCloudReposityImpl(*awsStorage)
	emailSendImpl, err := infrastruture.NewGomailSender("smtp.gmail.com")
	var sendEmailUsecase *user_usecase.SendAuthEmailUseCase
	if err != nil {
		fmt.Printf("Email service error: %v", err)
	}

	//userUsecases
	registerUserUsecase := user_usecase.NewRegisterUserUseCase(userRepoImpl)
	updateUserImage := user_usecase.NewUpdateUserImageUsecase(cloudStrImpl, userRepoImpl)
	loginUserUsecase := user_usecase.NewLoginUserUsecase(userRepoImpl)
	sendEmailUsecase = user_usecase.NewSendAuthEmailUseCase(userRepoImpl, emailSendImpl)
	verifyPasscodeUsecase := user_usecase.NewVerifyPasscodeUseCase(userRepoImpl)
	logoutUser := user_usecase.NewLogoutUserUseCase(userRepoImpl)

	//feedsUsecases
	postFeedusecase := feeds_usecase.NewCreateFieldUsecase(feedRepoImpl, cloudStrImpl)
	fetchFeedsUsecase := feeds_usecase.NewFetchFeedsUsecase(feedRepoImpl)

	userUsecases := user_usecase.NewUserCases(*registerUserUsecase, *updateUserImage, *loginUserUsecase, sendEmailUsecase, *verifyPasscodeUsecase, *logoutUser)
	feedsUsecases := feeds_usecase.NewFeedUsecases(*postFeedusecase, *fetchFeedsUsecase)

	//controllers objects
	userController := controllers.NewUserController(*userUsecases)
	feedController := controllers.NewFeedsController(*feedsUsecases, feedChan)

	authorizer := middlewares.NewAuthMiddleware(*userRepoImpl)

	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/update/profile-img", authorizer.Authenticate(userController.UpadateUserImage)).Methods("POST")
	r.HandleFunc("/login", userController.UserLogin).Methods("POST")
	r.HandleFunc("/send-auth-email", authorizer.Authenticate(userController.SendAuthEmail)).Methods("POST")
	r.HandleFunc("/verify-passcode", authorizer.Authenticate(userController.VerifyPasscode)).Methods("POST")
	r.HandleFunc("/logout", authorizer.Authenticate(userController.LogoutUser)).Methods("DELETE")

	r.HandleFunc("/feed", authorizer.Authenticate(feedController.CreateFeed)).Methods("POST")
	r.HandleFunc("/feeds", authorizer.Authenticate(feedController.FetchFeeds)).Methods("GET")
}
