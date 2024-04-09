package main

import (
	"ekoa-certificate-generator/pkg/curseduca"
	"ekoa-certificate-generator/pkg/utils"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerCloudWatchEvent(ev events.CloudWatchAlarmTrigger) error {
	username := utils.GetEnv("PROF_CURSEDUCA_USERNAME", "")
	password := utils.GetEnv("PROF_CURSEDUCA_PASSWORD", "")

	auth, err := curseduca.Login(username, password)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("INFO: token: %+v\n", auth.AccessToken)

	reports, err := curseduca.FindReportEnrollment(auth)
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