package bucket

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// SaveFile implements IBucket.
func (b *Bucket) SaveFile(body []byte, key string, bucketName string) error {
	_, err := b.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return err
	}

	return nil
}
