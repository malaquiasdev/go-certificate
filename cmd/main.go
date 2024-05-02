package main

import (
	"ekoa-certificate-generator/config"
	"ekoa-certificate-generator/internal/curseduca"
	"log"
)

func main() {
	c := config.LoadConfig(true)

	cur, err := curseduca.NewClient(c.Curseduca)
	if err != nil {
		log.Fatal("ERROR: failed to connect with curseduca", err)
	}

	res, err := cur.GetReportEnrollment(10)
	if err != nil {
		log.Fatal("ERROR: failed to GET reports", err)
	}

	log.Printf("INFO: res - %+v\n", res)

}
