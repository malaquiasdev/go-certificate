package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	"ekoa-certificate-generator/internal/queue"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerImporter(ev events.CloudWatchAlarmTrigger) error {
	c := config.LoadConfig(false)

	queue, err := queue.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with SQS", err)
		return err
	}

	db, err := db.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with DynamoDB", err)
		return err
	}

	auth, err := curseduca.Login(c.Curseduca)
	if err != nil {
		log.Fatal("ERROR: failed to connect with curseduca", err)
		return err
	}

	reports, err := curseduca.FindReportEnrollment(auth, c.Curseduca)
	if err != nil {
		log.Fatal("ERROR: failed to find report enrollment", err)
		return err
	}

	log.Printf("INFO: reports totalCount - %+v\n", reports.Metadata.TotalCount)

	count := 0
	for _, data := range reports.Data {
		blocked := strings.Contains(c.Curseduca.BlockList, fmt.Sprint(data.Content.ID))
		if blocked {
			log.Printf("WARN: skipping training course blocked | ContentID - %+v\n", data.Content.ID)
			continue
		}

		if data.FinishedAt == nil {
			log.Printf("WARN: skipping report FinishedAt not found | ReportId - %+v\n", data.ID)
			continue
		}

		cert := model.Certificate{
			ReportId: data.ID,
		}

		dbRes, _ := db.Query(cert.GetFilterReportId(), "reportId", c.AWS.DynamoTableName)

		if len(dbRes.Items) != 0 {
			log.Printf("WARN: skipping certificate found | ReportId - %+v\n", data.ID)
			continue
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		messageBody := string(jsonData)
		err = queue.Send(messageBody, c.AWS.GeneretorQueueUrl)
		if err != nil {
			return err
		}
		count++
	}

	log.Printf("INFO: event message count - %+v\n", count)

	return nil
}

func main() {
	lambda.Start(handlerImporter)
}
