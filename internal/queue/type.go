package queue

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

type IQueue interface {
	Send(messageBody string, queueUrl string) error
}

type Queue struct {
	client *sqs.SQS
}
