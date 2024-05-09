package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func handleGetCertificates(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c := config.LoadConfig(false)
	db, err := db.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with DynamoDB", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	filter := expression.Name("publicUrl").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	dbRes, err := db.ScanAll(condition, "", c.AWS.DynamoTableName)
	if err != nil {
		log.Fatal("ERROR: failed to connect ScanAll DynamoDB", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	certificates := CertificatesDTO{
		Count: dbRes.Count,
	}
	
	if len(dbRes.Items) > 0 {
		certificates.NextPageKey, _ = db.ToString(dbRes.LastEvaluatedKey)
		for _, dbItem := range dbRes.Items {
			cert, _ := model.ParseDynamoToCertificate(dbItem)
			certificates.Items = append(certificates.Items, CertificateDTO{
				ID:        cert.PK,
				ContentId: cert.ContentId,
				StudentId: cert.StudentId,
				URL:       cert.PublicUrl,
			})

		}
	}

	res, err := json.Marshal(certificates)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(res),
	}, nil
}

func main() {
	lambda.Start(handleGetCertificates)
}
