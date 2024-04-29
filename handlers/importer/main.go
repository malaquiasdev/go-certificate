package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/models"
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
		blocked := strings.Contains(c.Curseduca.BlockList, fmt.Sprint(data.Content.ID))
		if blocked {
			log.Printf("WARN: skipping training course blocked - %+v\n", data)
			continue
		}

		if data.FinishedAt == nil {
			log.Printf("WARN: skipping report FinishedAt not found - %+v\n", data)
			continue
		}

		log.Printf("INFO: report data - %+v\n", data)

		filter := models.Certificate{
			StudentEmail: data.Member.Email,
		}

		log.Printf("INFO: certificate filter - %+v\n", filter)
		log.Printf("INFO: GetFilterReportId - %+v\n", filter.GetFilterEmail())

		dbRes, _ := db.GetOne(filter.GetFilterEmail(), c.AWS.DynamoTableName)
		if len(dbRes.Item) != 0 {
			log.Printf("WARN: skipping certificate found - %+v\n", dbRes.Item)
			continue
		}

		log.Printf("WARN: skipping certificate not found - %+v\n", dbRes.Item)

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
