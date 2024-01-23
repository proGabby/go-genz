package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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

	createdUser, err := u.userUsecases.RegisterUser.Execute(user.Name, user.Email, user.Password)

	fmt.Printf("created user is %v\n", createdUser)
	if err != nil {
		fmt.Printf("error is %v", err)
		utils.HandleError(map[string]interface{}{
			"error": err,
		}, http.StatusBadRequest, w)
		return
	}

	userDto := &dto.UserResponse{
		Name:            createdUser.Name,
		Email:           createdUser.Email,
		ProfileImageUrl: createdUser.ProfileImageUrl,
		Id:              createdUser.Id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDto)
}

func (u *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	// Decode the JSON request body into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("Invalid request body %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	resMap, err := u.userUsecases.LoginUser.Execute(user.Email, user.Password)

	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("Invalid request body %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resMap)

}

func (u *UserController) UpadateUserImage(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*entity.User)

	if !ok {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("file parsing error: %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	file, handler, err := r.FormFile("image")

	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": err,
		}, http.StatusBadRequest, w)
		return
	}

	defer file.Close()

	ok, extension := isAllowedImageFile(handler.Filename)

	extensionName := strings.TrimPrefix(extension, ".")

	if !ok {
		utils.HandleError(map[string]interface{}{
			"error": "uploaded file is not an allowed image format",
		}, http.StatusBadRequest, w)
		return
	}

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": err,
		}, http.StatusInternalServerError, w)
		return
	}

	savePath := filepath.Join("uploads", handler.Filename)

	osfile, err := os.Create(savePath)

	if err != nil {
		fmt.Printf("err os %v", err)
		utils.HandleError(map[string]interface{}{
			"error": err,
		}, http.StatusBadRequest, w)
		return
	}

	_, err = io.Copy(osfile, file)
	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("file copying error: %v", err),
		}, http.StatusInternalServerError, w)
		return
	}

	now := time.Now().Unix()
	userRes, err := u.userUsecases.UpdateProfile.Execute(user.Id, fmt.Sprintf("%v-image-%v%v", user.Email, now, extension), &extensionName, osfile)
	if err != nil {
		fmt.Printf("err n %v", err)
		utils.HandleError(map[string]interface{}{
			"error": err,
		}, http.StatusBadRequest, w)
		return
	}

	osfile.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRes)

}

func isAllowedImageFile(filename string) (bool, string) {
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
		".webp": true,
	}

	ext := filepath.Ext(filename)
	return allowedExtensions[strings.ToLower(ext)], ext
}

func (u *UserController) SendAuthEmail(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*entity.User)

	if user == nil || !ok {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}

	passcode, err := utils.GeneratePasscode(6)

	emailBody := fmt.Sprintf("<p>Dear User,</p><p>Your passcode is: <strong>%s</strong></p><p>Use this passcode to complete your action. Please do not share it with others.</p>", passcode)

	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}

	errr := u.userUsecases.SendAuthEmail.Execute(user.Id, "gaby12645@gmail.com", user.Email, "Verify Your Account", emailBody, passcode)

	if errr != nil {
		utils.HandleError(map[string]interface{}{
			"error": "passcode sending error",
		}, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "email sent successfully",
	})
}

func (u *UserController) VerifyPasscode(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*entity.User)

	if user == nil || !ok {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}

	var passcode struct {
		Passcode string `json:"passcode"`
	}

	err := json.NewDecoder(r.Body).Decode(&passcode)
	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": "invalid request body",
		}, http.StatusBadRequest, w)
		return
	}

	userRes, err := u.userUsecases.VerifyPasscode.Execute(user.Id, passcode.Passcode)

	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("passcode verification error: %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg":  "passcode verified successfully",
		"data": *userRes,
	})
}
