package certificate

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Certificate struct {
	PK        string `json:"PK,omitempty"`
	SK        string `json:"SK,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func InterfaceToModel(data interface{}) (instance *Certificate, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Certificate) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"id": p.PK}
}

func (p *Certificate) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Certificate) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"id":   p.PK,
		"name": p.Name,
	}
}

func ParseDynamoAtributeToStruct(response map[string]*dynamodb.AttributeValue) (p Certificate, err error) {
	if len(response) == 0 {
		return p, errors.New("item not found")
	}
	for key, value := range response {
		if key == "id" {
			p.PK = *value.S
		}
		if key == "name" {
			p.Name = *value.S
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}
