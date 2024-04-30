package bucket

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetFile implements IBucket.
func (b *Bucket) GetFile(key string, bucketName string) (*s3.GetObjectOutput, error) {
	obj, err := b.client.GetObject(&s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// GetFileBytes implements IBucket.
func (b *Bucket) GetFileBytes(key string, bucketName string) ([]byte, error) {
	obj, err := b.client.GetObject(&s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}
	byt, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	return byt, nil
}
