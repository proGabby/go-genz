package infrastruture

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

type AzureCloudInfrasture struct {
	accountName   string
	containerName string
	accessKey     string
}

func NewAzureCloudInfrasture() (*AzureCloudInfrasture, error) {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")

	if !ok {
		fmt.Printf("AZure storage account name not found on env")
		return nil, fmt.Errorf("AZure storage account name not found on env")
	}

	accessKey, ok := os.LookupEnv("AZURE_STORAGE_ACCESS_KEY")
	if !ok {
		fmt.Printf("Azure storage access key not found on env")
		return nil, fmt.Errorf("Azure storage access key not found on env")
	}

	containerName, ok := os.LookupEnv("AZURE_STORAGE_CONTAINER_NAME")

	if !ok {
		fmt.Printf("Azure storage container name not found on env")
		return nil, fmt.Errorf("Azure storage container name not found on env")
	}

	return &AzureCloudInfrasture{
		accountName:   accountName,
		containerName: containerName,
		accessKey:     accessKey,
	}, nil
}

func (az *AzureCloudInfrasture) UploadImageToAzureStorage(fileName string, fileExtension *string, fileHandler *os.File) (*string, error) {

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", az.accountName)

	cred, err := azblob.NewSharedKeyCredential(az.accountName, az.accessKey)

	if err != nil {
		return nil, err
	}

	client, err := azblob.NewClientWithSharedKeyCredential(serviceURL, cred, nil)

	if err != nil {
		return nil, err
	}

	contentType := fmt.Sprintf("image/%s", *fileExtension)
	val, err := client.UploadFile(context.TODO(), az.containerName, fileName, fileHandler,
		&azblob.UploadFileOptions{
			BlockSize:   int64(1024),
			Concurrency: uint16(3),
			//set http content type
			HTTPHeaders: &blob.HTTPHeaders{
				BlobContentType: &contentType,
			},
			Progress: func(bytesTransferred int64) {
				fmt.Println(bytesTransferred)
			},
		})

	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return nil, err
	}

	fmt.Printf("The blob was uploaded successfully details is: %v.\n", val)
	imageUrl := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", az.accountName, az.containerName, fileName)

	return &imageUrl, nil

}
