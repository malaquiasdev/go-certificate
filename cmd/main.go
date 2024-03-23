package main

import (
	"ekoa-certificate-generator/pkg/curseduca"
	"ekoa-certificate-generator/pkg/utils"
	"log"
)

func main() {
	username := utils.GetEnv("PROF_CURSEDUCA_USERNAME", "")
	password := utils.GetEnv("PROF_CURSEDUCA_PASSWORD", "")

	auth, err := curseduca.Login(username, password)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Printf("INFO: token: %+v\n", auth.AccessToken)

	reports, err := curseduca.FindEnrollments(auth)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Printf("INFO: Reports details: %+v\n", reports)
}
