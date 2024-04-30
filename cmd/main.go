package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/db"
	"ekoa-certificate-generator/internal/db/models"
	"log"
)

func main() {
	c := config.LoadConfig(true)
	sess := config.CreateAWSSession(c.AWS)
	db := db.Init(sess)

	cert := models.Certificate{
		StudentEmail: "",
	}

	cond, err := cert.GetFilterEmail()
	if err != nil {
		log.Printf("WARN: skipping GetFilterEmail")
	}

	log.Printf("INFO: key -", cond)
	dbRes, err := db.Query(cond, "studentEmail", c.AWS.DynamoTableName)
	if err != nil {
		log.Printf("WARN: query error - %+v\n", err)
	}

	if len(dbRes.Items) != 0 {
		log.Printf("WARN: skipping certificate found - %+v\n", dbRes.Items)
	}

	log.Printf("WARN: skipping certificate not found - %+v\n", dbRes.Items)
}
