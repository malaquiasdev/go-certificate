package queue

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Send implements IQueue.
func (q *Queue) Send(messageBody string, queueUrl string) error {
	res, err := q.client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(messageBody),
	})
	if err != nil {
		log.Fatal("ERROR: failed to send sqs message", err)
		return err
	}

	log.Printf("INFO: sent sqs message - %+v\n", res.MessageId)
	return nil
}
