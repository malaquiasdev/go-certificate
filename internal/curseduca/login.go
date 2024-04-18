package curseduca

import (
	"bytes"
	"ekoa-certificate-generator/config"
	"encoding/json"
	"io"
	"net/http"
)

type Member struct {
	IsAdmin bool   `json:"isAdmin"`
	ID      int    `json:"id"`
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Tenant  Tenant `json:"tenant"`
}

type Tenant struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Slug string `json:"slug"`
}

type Auth struct {
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	RedirectUrl      string `json:"redirectUrl"`
	ExpiresAt        string `json:"expiresAt"`
	AuthenticationId int    `json:"authenticationId"`
	CurrentLoginId   string `json:"currentLoginId"`
	Member           Member `json:"member"`
}

func Login(config config.Curseduca) (Auth, error) {
	url := config.ProfBaseUrl + "/login"

	payload, _ := json.Marshal(map[string]string{
		"username": config.Username,
		"password": config.Password,
	})

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return Auth{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", config.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return Auth{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Auth{}, err
	}

	var response Auth
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Auth{}, err
	}

	return response, nil
}
