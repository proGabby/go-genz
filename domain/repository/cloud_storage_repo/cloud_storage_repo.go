package cloudstorage

import "os"

type CloudStorageRepo interface {
	UploadImageToCloudStorage(fileName string, fileHandler *os.File) (*string, error);
}