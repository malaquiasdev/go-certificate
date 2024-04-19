package main

import (
	"ekoa-certificate-generator/internal/curseduca"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerGenerator(ev events.SQSEvent) error {
	report := curseduca.Course{}
	
	err := json.Unmarshal([]byte(ev.Records[0].Body), &report)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Printf("INFO: %+v\n", report)

	return nil
}

func main() {
	lambda.Start(handlerGenerator)
}
