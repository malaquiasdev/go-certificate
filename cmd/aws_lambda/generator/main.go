package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/bucket"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/db/model"
	imgDraw "ekoa-certificate-generator/internal/image_draw"
	"ekoa-certificate-generator/internal/queue"
	"ekoa-certificate-generator/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerGenerator(ev events.SQSEvent) error {
	report := curseduca.Report{}

	c := config.LoadConfig(false)
	queue, err := queue.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with SQS", err)
		return err
	}

	b, err := bucket.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with Bucket S3", err)
		return err
	}

	err = json.Unmarshal([]byte(ev.Records[0].Body), &report)
	if err != nil {
		log.Fatal("ERROR: parse event body to report", err)
		return err
	}

	log.Printf("INFO: report - %+v\n", report)

	if report.FinishedAt == nil {
		log.Printf("WARN: skipping report FinishedAt not found")
		return nil
	}

	cert := model.Certificate{
		ReportId:          report.ID,
		ContentId:         report.Content.ID,
		ContentSlug:       report.Content.Slug,
		ContentTitle:      report.Content.Title,
		CourseFinishedAt:  *report.FinishedAt,
		StudentId:         report.Member.ID,
		StudentName:       report.Member.Name,
		StudentEmail:      report.Member.Email,
		ExpiresAt:         *report.ExpiresAt,
		ExpirationEnabled: report.ExpirationEnabled,
	}
	cert.GenerateID()
	cert.SetFilePath()
	cert.SetPublicUrl(c.UrlPrefix)

	coverPath := "pdf_templates/" + fmt.Sprintf("%d%s", report.Content.ID, "/page_1.PNG")
	backCoverPath := "pdf_templates/" + fmt.Sprintf("%d%s", report.Content.ID, "/page_2.PNG")

	coverImg, err := b.GetFileBytes(coverPath, c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET cover image", err)
		return err
	}
	backCoverImg, err := b.GetFileBytes(backCoverPath, c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET back cover image", err)
		return err
	}

	formattedFinishedAt, _ := utils.FormatDateTimeToDateOnly(report.FinishedAt)

	fontSans, err := b.GetFileBytes("pdf_templates/fonts/EncodeSansExpanded-Bold.ttf", c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET EncodeSansExpanded font", err)
		return err
	}

	fontMont, err := b.GetFileBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET Montserrat font", err)
		return err
	}

	fontSign, err := b.GetFileBytes("pdf_templates/fonts/Thesignature.ttf", c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET Thesignature font", err)
		return err
	}

	img, err := imgDraw.Png(coverImg, []imgDraw.DrawParams{{
		Key: "FULL_NAME",
		Text: imgDraw.FieldText{
			Position: imgDraw.Position{
				X: 610,
				Y: 430,
			},
			Font: imgDraw.Font{
				Size: 50.0,
				File: fontSans,
			},
			Value: report.Member.Name,
		},
	}, {
		Key: "FINISH_AT",
		Text: imgDraw.FieldText{
			Position: imgDraw.Position{
				X: 1350,
				Y: 665,
			},
			Font: imgDraw.Font{
				Size: 35.0,
				File: fontMont,
			},
			Value: formattedFinishedAt,
		},
	}, {
		Key: "SIGNATURE",
		Text: imgDraw.FieldText{
			Position: imgDraw.Position{
				X: 1300,
				Y: 830,
			},
			Font: imgDraw.Font{
				Size: 70.0,
				File: fontSign,
			},
			Value: strings.ToLower(report.Member.Name),
		},
	}, {
		Key: "AUTHENTITCATION_KEY",
		Text: imgDraw.FieldText{
			Position: imgDraw.Position{
				X: 500,
				Y: 1030,
			},
			Font: imgDraw.Font{
				Size: 20.0,
				File: fontMont,
			},
			Value: cert.PublicUrl,
		},
	}})
	if err != nil {
		log.Fatal("ERROR: failed to draw image", err)
		return err
	}

	pdf, err := imgDraw.ToPdf(img, backCoverImg)
	if err != nil {
		log.Fatal("ERROR: failed to convert image to PDF", err)
		return err
	}

	b.SaveFile(pdf.Bytes(), cert.FilePath, c.AWS.BucketName)

	certStr, err := cert.ToString()
	if err != nil {
		log.Fatal("ERROR: parse certificate to string", err)
		return err
	}

	queue.Send(certStr, c.AWS.IndexerQueueUrl)

	return nil
}

func main() {
	lambda.Start(handlerGenerator)
}
