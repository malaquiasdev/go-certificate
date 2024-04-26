package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/models"
	"ekoa-certificate-generator/internal/queue"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerImporter(ev events.CloudWatchAlarmTrigger) error {
	c := config.LoadConfig(false)
	sess := config.CreateAWSSession(c.AWS)
	db := db.Init(sess)

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

		filter := models.Certificate{
			ReportId: data.ID,
		}

		dbRes, _ := db.GetOne(filter.GetFilterReportId(), c.AWS.DynamoTableName)
		cert, err := models.ParseDynamoAtributeToStruct(dbRes.Item)
		if err == nil {
			log.Printf("WARN: skipping certificate found to ReportId - %+v\n", cert)
			continue
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		jsonString := string(jsonData)
		queue.SendMessage(string(jsonString), c.AWS.GeneretorQueueUrl, sess)
		count++
	}

	log.Printf("INFO: event message count - %+v\n", count)

	return nil
}

func main() {
	lambda.Start(handlerImporter)
}
