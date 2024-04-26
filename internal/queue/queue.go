package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessage(messageBody string, queueUrl string, awsSession *session.Session) {
	client := sqs.New(awsSession)
	client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(messageBody),
	})
}
