package curseduca

import (
	"bytes"
	"ekoa-certificate-generator/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
	"io"
)

type Member struct {
	IsAdmin  bool   `json:"isAdmin"`
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Tenant   Tenant `json:"tenant"`
}

type Tenant struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Slug string `json:"slug"`
}

type Auth struct {
	AccessToken   string `json:"accessToken"`
	RefreshToken  string `json:"refreshToken"`
	RedirectUrl  string `json:"redirectUrl"`
	ExpiresAt    string `json:"expiresAt"`
	AuthenticationId int    `json:"authenticationId"`
	CurrentLoginId string `json:"currentLoginId"`
	Member        Member  `json:"member"`
}

func Login(username string, password string) (Auth, error) {
	baseUrl := utils.GetEnv("PROF_CURSEDUCA_BASE_URL", "")
	apiKey := utils.GetEnv("CURSEDUCA_API_KEY", "")
	url := baseUrl + "/login"
	
	payload, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

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

	var response Auth
	err = json.Unmarshal(body, &response)
	if err != nil {
	  return Auth{}, err
	}
  
	return response, nil
}