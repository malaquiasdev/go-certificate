package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/utils"
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
	log.Printf("INFO: reports totalCount - %+v\n", reports.Metadata.TotalCount)

	count := 0
	for _, data := range reports.Data {
		if data.FinishedAt == nil {
			log.Printf("WARN: skipping report FinishedAt not found - %+v\n", data)
			continue
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		jsonString := string(jsonData)
		utils.QueueSendMessage(string(jsonString), c.AWS.GeneretorQueueUrl, sess)
		count++
	}

	log.Printf("INFO: event message count - %+v\n", count)

	return nil
}

func main() {
	lambda.Start(handlerImporter)
}
