package queue

import (
	"ekoa-certificate-generator/config"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewClient(c config.AWS) (IQueue, error) {
	sess, err := config.CreateAWSSession(c)
	if err != nil {
		return nil, err
	}

	client := sqs.New(sess)

	return &Queue{
		client: client,
	}, nil
}
