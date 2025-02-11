package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/model"
	"ekoa-certificate-generator/internal/queue"
	"ekoa-certificate-generator/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlerImporter(ev events.CloudWatchAlarmTrigger) error {
	c := config.LoadConfig(false)

	queue, err := queue.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with SQS", err)
		return err
	}

	dynamo, err := db.NewClient(c.AWS)
	if err != nil {
		log.Fatal("ERROR: failed to connect with DynamoDB", err)
		return err
	}

	mysql, err := db.NewMySQLClient(c.Mysql)
	if err != nil {
		log.Fatal("ERROR: failed to connect with MySQL", err)
	}

	cur, err := curseduca.NewClient(c.Curseduca)
	if err != nil {
		log.Fatal("ERROR: failed to connect with curseduca", err)
		return err
	}

	bilionLimit := 10000000000000000
	res, err := cur.GetReportEnrollment(bilionLimit)
	if err != nil {
		log.Fatal("ERROR: failed to find report enrollment", err)
		return err
	}

	log.Printf("INFO: totalCount - %+v\n", res.Metadata.TotalCount)

	count := 0
	for _, report := range res.Reports {
		blocked := strings.Contains(c.Curseduca.BlockList, fmt.Sprint(report.Content.ID))
		if blocked {
			log.Printf("WARN: skipping training course blocked | ContentID - %+v\n", report.Content.ID)
			continue
		}

		if report.FinishedAt == nil {
			log.Printf("WARN: skipping report FinishedAt not found | ReportId - %+v\n", report.ID)
			continue
		}

		cert := model.Certificate{
			ReportId: report.ID,
		}

		dbRes, _ := dynamo.Query(cert.GetFilterReportId(), "reportId", c.AWS.DynamoTableName)
		if len(dbRes.Items) != 0 {
			log.Printf("WARN: skipping certificate found | ReportId - %+v\n", report.ID)
			continue
		}
		course, _ := queryCourse(mysql, report.Content.ID)

		member, err := cur.GetMemberById(report.Member.ID)
		if err != nil {
			log.Fatal("ERROR: failed to get member", err)
			return err
		}

		report.Member.Document = member.Document
		report.ExpiresAt = utils.CalculateExpirationDate(report.FinishedAt, course.ValidationDays)

		jsonData, err := json.Marshal(report)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		messageBody := string(jsonData)
		err = queue.Send(messageBody, c.AWS.GeneretorQueueUrl)
		if err != nil {
			return err
		}
		count++
	}

	log.Printf("INFO: event message count - %+v\n", count)

	return nil
}

func queryCourse(mysql db.IMySQL, contentId int) (*model.Course, error) {
	rows, err := mysql.GetDB().Query(`
		SELECT id, idsCurseduca, validadeEmDias
		FROM railway.curso 
		WHERE JSON_EXTRACT(CAST(idsCurseduca AS JSON), '$.content_id') = ?
	`, contentId)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var course model.Course
	if !rows.Next() {
		return nil, fmt.Errorf("course not found for content_id: %d", contentId)
	}

	err = rows.Scan(
		&course.ID,
		&course.CurseducaIds,
		&course.ValidationDays,
	)
	if err != nil {
		return nil, fmt.Errorf("error scanning course: %v", err)
	}

	return &course, nil
}

func main() {
	lambda.Start(handlerImporter)
}
