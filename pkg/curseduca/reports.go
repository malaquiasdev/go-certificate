package curseduca

import (
	"ekoa-certificate-generator/pkg/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Metadata struct {
	TotalCount int `json:"totalCount"`
	HasMore    bool   `json:"hasmore"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

type Content struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type EnrollmentsMember struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Email     string   `json:"email"`
	GroupIds  []int    `json:"groupIds"`
}

type Course struct {
	ID             int       `json:"id"`
	Content        Content   `json:"content"`
	StartedAt      *string   `json:"startedAt"` // Pointer for handling null value
	FinishedAt     *string   `json:"finishedAt"`  // Pointer for handling null value
	Member         EnrollmentsMember	`json:"member"`
	SituationID    int       `json:"situationId"`
	Progress       int       `json:"progress"`
	ExpiresAt      string    `json:"expiresAt"`
	ExpirationEnabled bool     `json:"expirationEnabled"`
	Integration     string    `json:"integration"`
}

type ReportEnrollment struct {
	Metadata Metadata `json:"metadata"`
	Data      []Course `json:"data"`
}

func FindEnrollments(auth Auth) (ReportEnrollment, error) {
	baseUrl := utils.GetEnv("CLASS_CURSEDUCA_BASE_URL", "")
	apiKey := utils.GetEnv("CURSEDUCA_API_KEY", "")
	url := baseUrl + "/reports/enrollments?limit=1000000000"
	
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", apiKey)
	req.Header.Set("Authorization", "Bearer " + auth.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var response ReportEnrollment
	err = json.Unmarshal(body, &response)
	if err != nil {
	  return ReportEnrollment{}, err
	}
  
	return response, nil
}