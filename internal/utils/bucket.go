package utils

import (
	"bytes"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func BucketGetObject(key string, bucketName string, awsSession *session.Session) (*s3.GetObjectOutput, error) {
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

func BucketGetObjectBytes(key string, bucketName string, awsSession *session.Session) []byte {
	bucketObj, err := BucketGetObject(key, bucketName, awsSession)
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

func BucketSaveObject(body []byte, key string, bucketName string, awsSession *session.Session) error {
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
