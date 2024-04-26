package bucket

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetFile(key string, bucketName string, awsSession *session.Session) (*s3.GetObjectOutput, error) {
	client := s3.New(awsSession)

	obj, err := client.GetObject(&s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal("ERROR: BucketGetObject", err)
		return nil, err
	}

	return obj, nil
}

func GetFileBytes(key string, bucketName string, awsSession *session.Session) []byte {
	bucketObj, err := GetFile(key, bucketName, awsSession)
	if err != nil {
		log.Fatal("ERROR: failed get bucket image ", err)
		panic(err)
	}

	by, err := io.ReadAll(bucketObj.Body)
	if err != nil {
		log.Fatal("ERROR: failed to decode bucketObj ", err)
		panic(err)
	}

	return by
}
