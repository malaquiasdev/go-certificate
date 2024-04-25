package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/imagedraw"
	"ekoa-certificate-generator/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerGenerator(ev events.SQSEvent) error {
	report := curseduca.Course{}
	c := config.LoadConfig(false)
	sess := config.CreateAWSSession(c.AWS)

	err := json.Unmarshal([]byte(ev.Records[0].Body), &report)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Printf("INFO: report - %+v\n", report)

	if report.FinishedAt == nil {
		log.Printf("WARN: skipping report FinishedAt not found")
		return nil
	}

	imgPage1Path := "pdf_templates/" + fmt.Sprintf("%d%s", report.Content.ID, "/page_1.PNG")
	imgPage2Path := "pdf_templates/" + fmt.Sprintf("%d%s", report.Content.ID, "/page_2.PNG")

	imgPage1 := utils.BucketGetObjectBytes(imgPage1Path, c.AWS.BucketName, sess)
	imgPage2 := utils.BucketGetObjectBytes(imgPage2Path, c.AWS.BucketName, sess)

	formattedFinishedAt, _ := utils.FormatDateTimeToDateOnly(report.FinishedAt)

	imgDraw := imagedraw.DrawAndEconde(imgPage1, []imagedraw.Field{{
		Key: "FULL_NAME",
		Text: imagedraw.FieldText{
			FontSize:  50.0,
			PositionX: 610,
			PositionY: 430,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/EncodeSansExpanded-Bold.ttf", c.AWS.BucketName, sess),
			Value:     report.Member.Name,
		},
	}, {
		Key: "FINISH_AT",
		Text: imagedraw.FieldText{
			FontSize:  35.0,
			PositionX: 1350,
			PositionY: 665,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName, sess),
			Value:     formattedFinishedAt,
		},
	}, {
		Key: "SIGNATURE",
		Text: imagedraw.FieldText{
			FontSize:  70.0,
			PositionX: 1300,
			PositionY: 830,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Thesignature.ttf", c.AWS.BucketName, sess),
			Value:     strings.ToLower(report.Member.Name),
		},
	}})

	pdf := imagedraw.ImageToPdf(imgDraw, imgPage2)

	fileSavePath := "pdf/" + report.Member.Email + "/" + fmt.Sprintf("%d%s", report.ID, ".pdf")

	utils.BucketSaveObject(pdf.Bytes(), fileSavePath, c.AWS.BucketName, sess)

	return nil
}

func main() {
	lambda.Start(handlerGenerator)
}
