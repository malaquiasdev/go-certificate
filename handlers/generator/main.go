package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/imagedraw"
	"ekoa-certificate-generator/internal/utils"

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

	img := utils.BucketGetObjectBytes("pdf_templates/320/page_1.png", c.AWS.BucketName, sess)
	imageDraw := imagedraw.DrawAndEconde(img, []imagedraw.Field{{
		Key: "FULL_NAME",
		Text: imagedraw.FieldText{
			FontSize:  50.0,
			PositionX: 610,
			PositionY: 430,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName, sess),
			Value:     "Mateus Oliveira Malaquias",
		},
	}, {
		Key: "FINISH_AT",
		Text: imagedraw.FieldText{
			FontSize:  35.0,
			PositionX: 1350,
			PositionY: 665,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName, sess),
			Value:     "20/04/2024",
		},
	}, {
		Key: "SIGNATURE",
		Text: imagedraw.FieldText{
			FontSize:  40.0,
			PositionX: 1300,
			PositionY: 800,
			FontBytes: utils.BucketGetObjectBytes("pdf_templates/fonts/Allura-Regular.ttf", c.AWS.BucketName, sess),
			Value:     "Mateus Oliveira Malaquias",
		},
	}})

	// Upload modified image to S3 (replace with your upload function)
	utils.BucketSaveObject(imageDraw.Bytes(), "test.png", c.AWS.BucketName, sess)

	return nil
}

func main() {
	lambda.Start(handlerGenerator)
}
