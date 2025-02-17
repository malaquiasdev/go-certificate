package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/bucket"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	imgDraw "ekoa-certificate-generator/internal/image_draw"
	"ekoa-certificate-generator/internal/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var finishedAt *string = new(string)
	*finishedAt = "2023-09-17T15:36:10.000Z"

	c := config.LoadConfig(true)
	b, err := bucket.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with Bucket S3", err)
		return
	}

	db, err := db.NewMySQLClient(c.Mysql)
	if err != nil {
		log.Fatal("ERROR: failed to connect with MySQL", err)
	}

	rows, err := db.GetDB().Query(`
		SELECT id, idsCurseduca, validadeEmDias
		FROM railway.curso 
		WHERE JSON_EXTRACT(CAST(idsCurseduca AS JSON), '$.content_id') = ?
	`, 396)
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()

	var course model.Course
	if rows.Next() {
		err := rows.Scan(
			&course.ID,
			&course.CurseducaIds,
			&course.ValidationDays,
		)
		if err != nil {
			log.Fatal("Error scanning course:", err)
		}
		log.Printf("INFO: Course Details - %+v\n", course)
		log.Printf("INFO: Course Details - %+v\n", course.GetCurseducaIds().ContentId)
		log.Printf("INFO: ValidationDays - %d\n", *course.ValidationDays)
	}

	report := curseduca.Report{
		ID: 11796,
		Content: curseduca.Content{
			ID:    999999999999999999,
			Slug:  "gestao-de-tempo",
			Title: "Gestão de Tempo",
		},
		FinishedAt: finishedAt,
		Member: curseduca.EnrollmentsMember{
			ID:       6467,
			Name:     "Égnaldo Aécio Neves",
			Slug:     "egnaldo-aecio-001",
			Email:    "001@test.com",
			GroupIds: []int{40},
		},
		SituationID:       1,
		Progress:          100,
		ExpirationEnabled: true,
		Integration:       "Gratuito",
	}

	expiresAt := utils.CalculateExpirationDate(report.FinishedAt, course.ValidationDays)

	cert := model.Certificate{
		ReportId:          report.ID,
		ContentId:         report.Content.ID,
		ContentSlug:       report.Content.Slug,
		ContentTitle:      report.Content.Title,
		CourseFinishedAt:  *report.FinishedAt,
		StudentId:         report.Member.ID,
		StudentName:       report.Member.Name,
		StudentEmail:      report.Member.Email,
		ExpiresAt:         expiresAt,
		ExpirationEnabled: report.ExpirationEnabled,
	}
	cert.GenerateID()
	cert.SetFilePath()
	cert.SetPublicUrl("")

	cur, err := curseduca.NewClient(c.Curseduca)
	if err != nil {
		log.Fatal("ERROR: failed to connect with curseduca", err)
		return
	}

	member, err := cur.GetMemberById(cert.StudentId)
	if err != nil {
		log.Fatal("ERROR: failed to get member", err)
		return
	}

	coverPath := "pdf_templates/" + fmt.Sprintf("%d%s", report.Content.ID, "/page_1.PNG")
	backCoverPath := "pdf_templates/" + fmt.Sprintf("%d%s", report.Content.ID, "/page_2.PNG")

	coverImg, err := b.GetFileBytes(coverPath, c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET cover image", err)
		return
	}

	backCoverImg, err := b.GetFileBytes(backCoverPath, c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET back cover image", err)
		return
	}

	formattedFinishedAt, _ := utils.FormatDateTimeToDateOnly(report.FinishedAt)

	fontSans, err := b.GetFileBytes("pdf_templates/fonts/EncodeSansExpanded-Bold.ttf", c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET EncodeSansExpanded font", err)
		return
	}

	fontMont, err := b.GetFileBytes("pdf_templates/fonts/Montserrat-Regular.ttf", c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET Montserrat font", err)
		return
	}

	fontSign, err := b.GetFileBytes("pdf_templates/fonts/Thesignature.ttf", c.AWS.BucketName)
	if err != nil {
		log.Fatal("ERROR: GET Thesignature font", err)
		return
	}

	rsaKey, certificate, hash, err := b.GetPkcs(c.AWS.Pkcs.FileKey, c.AWS.BucketName, c.AWS.Pkcs.Password)
	if err != nil {
		log.Fatal("ERROR: GET certificate", err)
		return
	}

	log.Println(rsaKey)
	log.Println(certificate)
	log.Println(hash)

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
		Key: "EXPIRES_AT",
		Text: imgDraw.FieldText{
			Position: imgDraw.Position{
				X: 960,
				Y: 720,
			},
			Font: imgDraw.Font{
				Size: 35.0,
				File: fontMont,
			},
			Value: func() string {
				if expiresAt == "" {
					return ""
				}
				return "Data de expiração:       " + expiresAt
			}(),
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
			Value: strings.ToLower(utils.NormalizeString(report.Member.Name)),
		},
	}, {
		Key: "DOCUMENT",
		Text: imgDraw.FieldText{
			Position: imgDraw.Position{
				X: 1370,
				Y: 900,
			},
			Font: imgDraw.Font{
				Size: 15.0,
				File: fontSans,
			},
			Value: func() string {
				if member.Document == "" {
					return ""
				}
				return "CPF: " + member.Document
			}(),
		},
	},
		{
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
		return
	}

	pdf, err := imgDraw.ToPdf(img, backCoverImg)
	if err != nil {
		log.Fatal("ERROR: failed to convert image to PDF", err)
		return
	}

	fileName := fmt.Sprintf("certificate_%d.pdf", report.ID)
	err = os.WriteFile(fileName, pdf.Bytes(), 0644)
	if err != nil {
		log.Fatal("ERROR: failed to save PDF locally", err)
		return
	}
	log.Printf("INFO: PDF saved successfully as %s", fileName)
}
