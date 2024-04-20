package main

import (
	"bytes"
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/certificate"
	"ekoa-certificate-generator/internal/utils"
	"image/png"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerGenerator(ev events.SQSEvent) error {
	// report := curseduca.Course{}
	c := config.LoadConfig(false)
	sess := config.CreateAWSSession(c.AWS)

	/*err := json.Unmarshal([]byte(ev.Records[0].Body), &report)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}*/
	// log.Printf("INFO: report %q", report)

	// Decode image
	img := utils.BucketGetObjectBytes("pdf_templates/320/page_1.png", c.AWS.BucketName, sess)

	fullName := certificate.Field{
		Key: "FULL_NAME",
		Text: certificate.FieldText{
			FontSize:  50.0,
			PositionX: 610,
			PositionY: 430,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName, sess),
			Value:     "Mateus Oliveira Malaquias",
		},
	}

	finishAt := certificate.Field{
		Key: "FINISH_AT",
		Text: certificate.FieldText{
			FontSize:  35.0,
			PositionX: 1350,
			PositionY: 665,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName, sess),
			Value:     "20/04/2024",
		},
	}

	signature := certificate.Field{
		Key: "SIGNATURE",
		Text: certificate.FieldText{
			FontSize:  40.0,
			PositionX: 1300,
			PositionY: 800,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Allura-Regular.ttf", c.AWS.BucketName, sess),
			Value:     "Mateus Oliveira Malaquias",
		},
	}
	imageDraw, _ := certificate.Generate(img, []certificate.Field{fullName, finishAt, signature})

	b := new(bytes.Buffer)
	if err := png.Encode(b, imageDraw); err != nil {
		log.Fatal("ERROR: unable to encode image ", err)
		panic(err)
	}

	// Upload modified image to S3 (replace with your upload function)
	utils.BucketSaveObject(b.Bytes(), "test.png", c.AWS.BucketName, sess)

	return nil
}

func main() {
	lambda.Start(handlerGenerator)
}
