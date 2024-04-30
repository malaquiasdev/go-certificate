package bucket

import (
	"ekoa-certificate-generator/config"

	"github.com/aws/aws-sdk-go/service/s3"
)

func NewClient(c config.AWS) (IBucket, error) {
	sess, err := config.CreateAWSSession(c)
	if err != nil {
		return nil, err
	}

	client := s3.New(sess)
	return &Bucket{
		client: client,
	}, nil
}
