package cloudstorage

import "os"

type CloudStorageRepo interface {
	UploadImageToCloudStorage(fileName string, fileExtensionName *string, fileHandler *os.File) (*string, error)
}
