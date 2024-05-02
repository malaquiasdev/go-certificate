package curseduca

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (c *Curseduca) GetReportEnrollment(limit int) (ReportEnrollment, error) {
	url := c.httpConfig.classBaseUrl + "/reports/enrollments?" + "limit=" + strconv.Itoa(limit)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReportEnrollment{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", c.httpConfig.apiKey)
	req.Header.Set("Authorization", "Bearer "+c.httpConfig.auth.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return ReportEnrollment{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ReportEnrollment{}, err
	}

	var response ReportEnrollment
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ReportEnrollment{}, err
	}

	return response, nil
}
