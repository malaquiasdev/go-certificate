package curseduca

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (c *Curseduca) GetReportEnrollment(limit int) (ReportEnrollment, error) {
	startDate := time.Now().AddDate(0, 0, -90).Format("2006-01-02")
	url := c.httpConfig.classBaseUrl + "/enrollments?" + "limit=" + strconv.Itoa(limit) + "&startDate=" + startDate

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
