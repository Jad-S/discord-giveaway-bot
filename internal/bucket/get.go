package bucket

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"os"
)

// GetFile retrieves a file from S3 and returns the contents as a string.
func GetFile(filename string) ([]byte, error) {
	log.Println("Downloading: ", filename)

	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Fatalf("Error retrieving file from S3: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	log.Printf("Successfully retrieved: %s", filename)

	return body, nil
}
