package cloud_repo_impl

import (
	"os"

	infrastruture "github.com/proGabby/4genz/data/cloud_infrastruture"
)

type AZCloudReposityImpl struct {
	azcloudStrg infrastruture.AzureCloudInfrasture
}

type AwsCloudReposityImpl struct {
	awscloudStrg infrastruture.S3Handler
}

func NewAWSCloudReposityImpl(awscloudStrg infrastruture.S3Handler) *AwsCloudReposityImpl {

	return &AwsCloudReposityImpl{
		awscloudStrg: awscloudStrg,
	}
}

func (cp *AwsCloudReposityImpl) UploadImageToCloudStorage(fileName string, fileExtensionName *string, fileHandler *os.File) (*string, error) {
	return cp.awscloudStrg.UploadFile(fileName, *fileExtensionName, fileHandler)
}
