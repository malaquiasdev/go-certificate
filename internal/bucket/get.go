package bucket

import (
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"software.sslmate.com/src/go-pkcs12"
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

func (b *Bucket) GetPkcs(key string, bucketName string, password string) (*rsa.PrivateKey, []byte, []byte, error) {
	obj, err := b.client.GetObject(&s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get pfx from S3: %v", err)
	}
	defer obj.Body.Close()

	pfxData, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read pfx data: %v", err)
	}

	// Decode the PKCS12 file
	privateKey, certificate, err := pkcs12.Decode(pfxData, password)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode pfx: %v", err)
	}

	// Type assert to get RSA private key
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, nil, fmt.Errorf("private key is not RSA")
	}

	// Generate SHA256 hash of the certificate
	hash := sha256.Sum256(certificate.Raw)

	return rsaKey, certificate.Raw, hash[:], nil
}
