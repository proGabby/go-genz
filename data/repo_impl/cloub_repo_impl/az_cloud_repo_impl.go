package cloud_repo_impl

import (
	"os"

	infrastruture "github.com/proGabby/4genz/data/cloud_infrastruture"
)

type CloudReposityImpl struct {
	azcloudStrg infrastruture.AzureCloudInfrasture
}

func NewCloudReposityImpl(azcloudStrg infrastruture.AzureCloudInfrasture) *CloudReposityImpl {

	return &CloudReposityImpl{
		azcloudStrg: azcloudStrg,
	}
}

func (az *CloudReposityImpl) UploadImageToCloudStorage(fileName string, fileExtensionName *string, fileHandler *os.File) (*string, error) {
	return az.azcloudStrg.UploadImageToAzureStorage(fileName, fileExtensionName, fileHandler)
}
