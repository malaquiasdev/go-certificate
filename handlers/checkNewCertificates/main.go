package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/pkg/curseduca"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerCloudWatchEvent(ev events.CloudWatchAlarmTrigger) error {
	config := config.LoadConfig(false)

	auth, err := curseduca.Login(config.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("INFO: token: %+v\n", auth.AccessToken)

	reports, err := curseduca.FindReportEnrollment(auth, config.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("INFO: Reports details: %+v\n", reports)

	return nil
}

func main() {
	lambda.Start(handlerCloudWatchEvent)
}
