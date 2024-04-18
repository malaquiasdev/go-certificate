package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/pkg/curseduca"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerImporter(ev events.CloudWatchAlarmTrigger) error {
	config := config.LoadConfig(false)

	auth, err := curseduca.Login(config.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	reports, err := curseduca.FindReportEnrollment(auth, config.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("INFO: Reports details: %+v\n", reports)

	return nil
}

func main() {
	lambda.Start(handlerImporter)
}
