package bucket

import (
	"crypto/rsa"

	"github.com/aws/aws-sdk-go/service/s3"
)

type IBucket interface {
	GetFile(key string, bucketName string) (*s3.GetObjectOutput, error)
	GetFileBytes(key string, bucketName string) ([]byte, error)
	SaveFile(body []byte, key string, bucketName string) error
	GetPkcs(key string, bucketName string, password string) (*rsa.PrivateKey, []byte, []byte, error)
}

type Bucket struct {
	client *s3.S3
}
