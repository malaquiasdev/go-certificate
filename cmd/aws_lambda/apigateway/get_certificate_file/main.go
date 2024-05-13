package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleGetCertificateFile(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(req)
	log.Println(req.PathParameters)
	c := config.LoadConfig(false)
	db, err := db.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with DynamoDB", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	cert := model.Certificate{
		PK: req.PathParameters["id"],
	}

	dbRes, err := db.GetOne(cert.GetFilterPK(), c.AWS.DynamoTableName)
	if err != nil {
		log.Fatal("ERROR: failed to get certificate", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	cert, err = model.ParseDynamoToCertificate(dbRes.Item)
	if err != nil {
		log.Fatal("ERROR: failed to parse db to model", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	log.Println(cert.FilePath)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/pdf",
		},
		Body:            "",
		IsBase64Encoded: true,
	}, nil
}

func main() {
	lambda.Start(handleGetCertificateFile)
}
