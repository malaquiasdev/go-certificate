package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/queue"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerImporter(ev events.CloudWatchAlarmTrigger) error {
	c := config.LoadConfig(false)
	sess := config.CreateAWSSession(c.AWS)

	auth, err := curseduca.Login(c.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	reports, err := curseduca.FindReportEnrollment(auth, c.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("INFO: Reports details: %+v\n", reports)

	for _, data := range reports.Data {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		
		jsonString := string(jsonData)
		queue.Send(string(jsonString), c.AWS.GeneretorQueueUrl, sess)
	}


	return nil
}

func main() {
	lambda.Start(handlerImporter)
}
