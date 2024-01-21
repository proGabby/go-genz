package user_usecase

import (
	"os"

	"github.com/proGabby/4genz/data/dto"
	cloudstorage "github.com/proGabby/4genz/domain/repository/cloud_storage_repo"
	"github.com/proGabby/4genz/domain/repository/user_repo"
)

type UpdateUserImageUsecase struct {
	cloudRepo cloudstorage.CloudStorageRepo
	userRepo  user_repo.UserRepository
}

func NewUpdateUserImageUsecase(cloudRepo cloudstorage.CloudStorageRepo, userRepo user_repo.UserRepository) *UpdateUserImageUsecase {

	return &UpdateUserImageUsecase{
		cloudRepo: cloudRepo,
		userRepo:  userRepo,
	}
}

func (c *UpdateUserImageUsecase) Execute(userId int, fileName string, fileExtensionName *string, fileHandler *os.File) (*dto.UserResponse, error) {

	imageUrl, err := c.cloudRepo.UploadImageToCloudStorage(fileName, fileExtensionName, fileHandler)

	if err != nil {
		return nil, err
	}

	userRes, err := c.userRepo.UpdateProfileImage(userId, *imageUrl)

	if err != nil {
		return nil, err
	}

	return userRes, nil

}
