package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/models"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerIndexer(ev events.SQSEvent) error {
	cert := models.Certificate{}

	c := config.LoadConfig(false)
	sess, _ := config.CreateAWSSession(c.AWS)
	db := db.Init(sess)

	err := json.Unmarshal([]byte(ev.Records[0].Body), &cert)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	cert.SetCreatedAt()
	cert.SetUpdatedAt()

	item, err := db.CreateOrUpdate(cert, c.AWS.DynamoTableName)
	if err != nil {
		log.Fatal("ERROR: CreateOrUpdate", err)
		panic(err)
	}
	log.Printf("INFO: item - %+v\n", item)

	return nil
}

func main() {
	lambda.Start(handlerIndexer)
}
