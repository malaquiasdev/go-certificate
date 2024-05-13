package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/db"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleGetCertificateFile(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(req)
	c := config.LoadConfig(false)
	_, err := db.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with DynamoDB", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	id := req.PathParameters["id"]

	log.Println(id)
	/*
	  cert := model.Certificate{
	    PK: string,
	  }
	*/

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
