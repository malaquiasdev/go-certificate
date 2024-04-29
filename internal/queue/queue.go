package queue

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessage(messageBody string, queueUrl string, awsSession *session.Session) {
	log.Printf("INFO: sending event message - %+v\n", messageBody)

	client := sqs.New(awsSession)
	_, err := client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(messageBody),
	})

	if err != nil {
		log.Fatal("ERROR: failed to send sqs message", err)
		panic(err)
	}
}
