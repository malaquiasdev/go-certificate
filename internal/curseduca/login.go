package curseduca

import (
	"bytes"
	"ekoa-certificate-generator/config"
	"encoding/json"
	"io"
	"net/http"
)

func login(config config.Curseduca) (auth, error) {
	url := config.ProfBaseUrl + "/login"

	payload, _ := json.Marshal(map[string]string{
		"username": config.Username,
		"password": config.Password,
	})

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return auth{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", config.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return auth{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return auth{}, err
	}

	var response auth
	err = json.Unmarshal(body, &response)
	if err != nil {
		return auth{}, err
	}

	return response, nil
}
