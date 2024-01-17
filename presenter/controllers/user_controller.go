package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/entity"
	"github.com/proGabby/4genz/domain/usecase/user_usecase"
	"github.com/proGabby/4genz/utils"
)

type UserController struct {
	userUsecases user_usecase.UserUseCases
	// authMiddleware
}

func NewUserController(usersUsecase user_usecase.UserUseCases) *UserController {
	return &UserController{
		userUsecases: usersUsecase,
	}
}

var validate = validator.New()

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user entity.User

	// Decode the JSON request body into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("Invalid request body %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		// Handle validation error
		errors := err.(validator.ValidationErrors)
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("Validation failed: %v", errors),
		}, http.StatusBadRequest, w)

		return
	}

	createdUser, err := u.userUsecases.RegisterUser.Execute(user.Name, user.Email, user.Password, user.ProfileImageUrl)

	fmt.Printf("created user is %v\n", createdUser)
	if err != nil {
		fmt.Printf("error is %v", err)
		utils.HandleError(map[string]interface{}{
			"error": err,
		}, http.StatusBadRequest, w)
		return
	}

	userDto := &dto.CreateUserResponse{
		Name:            createdUser.Name,
		Email:           createdUser.Email,
		ProfileImageUrl: createdUser.ProfileImageUrl,
		Id:              createdUser.Id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDto)
}
