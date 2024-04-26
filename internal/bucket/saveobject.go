package bucket

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func SaveFile(body []byte, key string, bucketName string, awsSession *session.Session) error {
	client := s3.New(awsSession)

	_, err := client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		log.Fatal("ERROR: failed to upload object to S3", err)
		return err
	}

	return nil
}
