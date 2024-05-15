package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	"ekoa-certificate-generator/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func handleGetCertificates(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(req.MultiValueQueryStringParameters["email"])
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
	if value, exist := req.MultiValueQueryStringParameters["email"]; exist {
		filter.And(expression.Name("studentEmail").Equal(expression.Value(value[0])))
	}

	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	dbRes, err := db.ScanAll(condition, c.AWS.DynamoTableName)
	if err != nil {
		log.Fatal("ERROR: failed to connect ScanAll DynamoDB", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	certificates := CertificatesDTO{
		Count: len(dbRes.Items),
	}

	if len(dbRes.Items) > 0 {
		for _, dbItem := range dbRes.Items {
			cert, _ := model.ParseDynamoToCertificate(dbItem)
			createdAtFormatted, _ := utils.FormatDateTimeToDateOnly(&cert.CreatedAt)
			finishedAtFormatted, _ := utils.FormatDateTimeToDateOnly(&cert.CourseFinishedAt)
			expiresAtFormatted, _ := utils.FormatDateTimeToDateOnly(&cert.ExpiresAt)
			certificates.Items = append(certificates.Items, CertificateDTO{
				ID:         cert.PK,
				ContentId:  cert.ContentId,
				StudentId:  cert.StudentId,
				CreatedAt:  createdAtFormatted,
				FinishedAt: finishedAtFormatted,
				ExpiresAt:  expiresAtFormatted,
				URL:        cert.PublicUrl,
			})

		}
	}

	res, err := json.Marshal(certificates)
	if err != nil {
		log.Fatal("ERROR: failed to parse certificates", err)
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
