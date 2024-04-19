package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateAWSSession(config AWS) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	})
	if err != nil {
		log.Printf("ERROR: Error creating session: %+v\n", err)
		panic(err)
	}

	return sess
}
