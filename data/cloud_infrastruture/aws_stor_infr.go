package infrastruture

import (
	"fmt"
	//"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3_REGION = "us-east-1"
	S3_BUCKET = "genz-blob"
	S3_ACL    = "public-read"
)

type S3Handler struct {
	Session *session.Session
	Bucket  string
}

func NewS3Handler() (*S3Handler, error) {

	awsAccesskeyId, ok := os.LookupEnv("AWS_ACCESS_KEY_ID")

	if !ok {
		fmt.Println("AWS_ACCESS_KEY_ID not found on env file")
		return nil, fmt.Errorf("AWS_ACCESS_KEY_ID not found on env file")
	}

	awsSecretKey, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY")

	if !ok {
		fmt.Println("AWS_SECRET_ACCESS_KEY not found on env file")
		return nil, fmt.Errorf("AWS_SECRET_ACCESS_KEY not found on env file")
	}

	return &S3Handler{
		Session: session.Must(session.NewSession(&aws.Config{
			Region: aws.String(S3_REGION),
			Credentials: credentials.NewStaticCredentials(
				awsAccesskeyId, awsSecretKey, "",
			),
		})),
		Bucket: S3_BUCKET,
	}, nil
}

func (h S3Handler) UploadFile(filename, extensionName string, file *os.File) (*string, error) {

	if file == nil {
		return nil, fmt.Errorf("file not found")
	}

	_, err := s3.New(h.Session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(filename),
		ACL:    aws.String(S3_ACL),
		Body:   file,
		ContentType: aws.String(fmt.Sprintf("image/%s", extensionName)),

	})


	if err != nil {
		fmt.Printf("Unable to upload %q to %q, %v", filename, h.Bucket, err)
		return nil, err
	}

	publicURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", h.Bucket, filename)

	fmt.Printf("The image uploaded successfully. %s", publicURL)

	return &publicURL, err
}

