package curseduca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Curseduca) GetMemberById(memberId int) (Member, error) {
	url := c.httpConfig.profBaseUrl + "/members/" + url.QueryEscape(fmt.Sprintf("%d", memberId))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Member{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", c.httpConfig.apiKey)
	req.Header.Set("Authorization", "Bearer "+c.httpConfig.auth.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return Member{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Member{}, err
	}

	var member Member
	err = json.Unmarshal(body, &member)
	if err != nil {
		return Member{}, err
	}

	return member, nil
}
