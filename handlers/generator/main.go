package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerGenerator(ev events.SQSEvent) error {
	for _, message := range ev.Records {
		log.Printf("INFO: %+v\n", message.Body)
	}

	return nil
}

func main() {
	lambda.Start(handlerGenerator)
}
