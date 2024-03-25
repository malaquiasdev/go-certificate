package main

import (
	"ekoa-certificate-generator/pkg/certificate"
	"ekoa-certificate-generator/pkg/curseduca"
	"ekoa-certificate-generator/pkg/utils"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	
	username := utils.GetEnv("PROF_CURSEDUCA_USERNAME", "")
	password := utils.GetEnv("PROF_CURSEDUCA_PASSWORD", "")

	auth, err := curseduca.Login(username, password)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// log.Printf("INFO: token: %+v\n", auth.AccessToken)

	reports, err := curseduca.FindReportEnrollment(auth)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// log.Printf("INFO: Reports details: %+v\n", reports)

	report := reports.Data[30]
	
	certificateImage, err := certificate.Generate("assets/certificate_front.jpg", certificate.Field{
		Key: "NOME_COMPLETO",
		Text: certificate.FieldText{
			FontSize: 100.0,
			PositionX: 120,
			PositionY: 450,
			FontPath: "assets/EncodeSansExpanded-Bold.ttf",
			Value: report.Member.Name,
		},
	})
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
