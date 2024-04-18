package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/certificate"
	"ekoa-certificate-generator/internal/curseduca"
	"image/jpeg"
	"log"
	"os"
	"time"
)

func main() {

	config := config.LoadConfig(true)

	auth, err := curseduca.Login(config.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// log.Printf("INFO: token: %+v\n", auth.AccessToken)

	reports, err := curseduca.FindReportEnrollment(auth, config.Curseduca)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// log.Printf("INFO: Reports details: %+v\n", reports)

	report := reports.Data[3]

	fullName := certificate.Field{
		Key: "FULL_NAME",
		Text: certificate.FieldText{
			FontSize:  50.0,
			PositionX: 120,
			PositionY: 450,
			FontPath:  "assets/fonts/Montserrat-Regular.ttf",
			Value:     report.Member.Name,
		},
	}

	contentName := certificate.Field{
		Key: "CONTENT_NAME",
		Text: certificate.FieldText{
			FontSize:  30.0,
			PositionX: 120,
			PositionY: 630,
			FontPath:  "assets/fonts/Montserrat-Regular.ttf",
			Value:     report.Content.Title,
		},
	}

	finishAt := certificate.Field{
		Key: "FINISH_AT",
		Text: certificate.FieldText{
			FontSize:  35.0,
			PositionX: 920,
			PositionY: 745,
			FontPath:  "assets/fonts/Montserrat-Regular.ttf",
			Value:     getFinishedAt(report),
		},
	}

	signature := certificate.Field{
		Key: "SIGNATURE",
		Text: certificate.FieldText{
			FontSize:  35.0,
			PositionX: 1000,
			PositionY: 850,
			FontPath:  "assets/fonts/Allura-Regular.ttf",
			Value:     report.Member.Name,
		},
	}
	certificateImage, err := certificate.Generate("assets/certificate_front.jpg", []certificate.Field{fullName, contentName, finishAt, signature})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	file, err := os.Create("outimage.jpeg")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()

	err = jpeg.Encode(file, certificateImage, nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("Image saved successfully!")
}

func getFinishedAt(report curseduca.Course) string {
	if report.FinishedAt == nil {
		now := time.Now()
		return now.Format("2006-01-02")
	}
	return *report.FinishedAt
}
