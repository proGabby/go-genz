package feeds_usecase

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/proGabby/4genz/domain/entity"
	cloudstorage "github.com/proGabby/4genz/domain/repository/cloud_storage_repo"
	"github.com/proGabby/4genz/domain/repository/feed_repo"
	"github.com/proGabby/4genz/utils"
)

type CreateFeedUsecase struct {
	feedRepo  feed_repo.FeedRepository
	cloudRepo cloudstorage.CloudStorageRepo
}

func NewCreateFieldUsecase(feedRepo feed_repo.FeedRepository, cloudRepo cloudstorage.CloudStorageRepo) *CreateFeedUsecase {

	return &CreateFeedUsecase{
		feedRepo:  feedRepo,
		cloudRepo: cloudRepo,
	}
}

func (cf *CreateFeedUsecase) Execute(userId int, caption string, files []*multipart.FileHeader) (*entity.Feed, error) {
	feed, err := cf.feedRepo.CreateNewFeed(userId, caption)

	if err != nil {
		return nil, err
	}

	resultCh := make(chan *string)
	errCh := make(chan error)

	for i, fileHandler := range files {
		num := i + 1

		go func(fileHandler *multipart.FileHeader) {
			osfile, extensionName, err := cf.setupFile(fileHandler)
			defer osfile.Close()

			if err != nil {
				fmt.Printf("error on fileHandler")
				errCh <- fmt.Errorf("error on %v filehandler %v", num, err)
				return
			}

			if osfile == nil || extensionName == nil {
				fmt.Printf("error on osfile")
				errCh <- fmt.Errorf("invalid file on iteration %v", num)
				return
			}
			now := time.Now().Unix()
			imageUrl, err := cf.cloudRepo.UploadImageToCloudStorage(fmt.Sprintf("%d-feedimage-%v%v", feed.Id, now, fileHandler.Filename), extensionName, osfile)
			if err != nil {
				fmt.Printf("error on file uploading")
				errCh <- fmt.Errorf("error on %v filehandler %v", num, err)
				return
			}
			resultCh <- imageUrl

		}(fileHandler)

	}

	for range files {
		select {
		case imageUrl := <-resultCh:
			if imageUrl != nil {
				feed, err = cf.feedRepo.AddFeedImage(feed, *imageUrl)
				if err != nil {
					fmt.Printf("error adding image to feed: %v\n", err)
				}
			}
		case err := <-errCh:
			fmt.Printf("error adding image to feed: %v\n", err)
		}
	}

	return feed, nil
}

// setupFile is a function that handles the file upload
// it returns the file, the extension name and an error if any
func (cf *CreateFeedUsecase) setupFile(fileHandler *multipart.FileHeader) (*os.File, *string, error) {
	file, err := fileHandler.Open()

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()

	ok, extension := utils.IsAllowedImageFile(fileHandler.Filename)

	extensionName := strings.TrimPrefix(extension, ".")

	if !ok {
		return nil, nil, fmt.Errorf("invalid image file format")
	}

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return nil, nil, err

	}

	savePath := filepath.Join("uploads", fileHandler.Filename)

	osfile, err := os.Create(savePath)

	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(osfile, file)

	if err != nil {
		return nil, nil, err
	}

	_, err = osfile.Seek(0, 0)

	if err != nil {
		return nil, nil, err
	}

	return osfile, &extensionName, nil

}
