package bucket

import "github.com/aws/aws-sdk-go/service/s3"

type IBucket interface {
	GetFile(key string, bucketName string) (*s3.GetObjectOutput, error)
	GetFileBytes(key string, bucketName string) ([]byte, error)
	SaveFile(body []byte, key string, bucketName string) error
}

type Bucket struct {
	client *s3.S3
}
