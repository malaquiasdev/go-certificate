package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/bucket"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	"encoding/base64"
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

	b, err := bucket.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with Bucket S3", err)
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

	pdfByte, err := b.GetFileBytes(cert.FilePath, c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET pdf file ", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}
	res := base64.StdEncoding.EncodeToString(pdfByte)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/pdf",
		},
		Body:            res,
		IsBase64Encoded: true,
	}, nil
}

func main() {
	lambda.Start(handleGetCertificateFile)
}
