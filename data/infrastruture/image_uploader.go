package infrastruture

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func UploadImageToBlobStorage(fileName string, fileHandler *os.File) (*string, error) {

	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		return nil, err
	}

	client, err := azblob.NewClient(serviceURL, cred, nil)

	if err != nil {
		return nil, err
	}

	val, err := client.UploadFile(context.TODO(), "testcontainer", "virtual/dir/path/"+fileName, fileHandler,
		&azblob.UploadFileOptions{
			BlockSize:   int64(1024),
			Concurrency: uint16(3),
			// If Progress is non-nil, this function is called periodically as bytes are uploaded.
			Progress: func(bytesTransferred int64) {
				fmt.Println(bytesTransferred)
			},
		})

	if err != nil {
		return nil, err
	}

	fmt.Printf("The blob was uploaded successfully details is: %v.\n", val)

	res := "empty"

	return &res, nil

}
