package models

import (
	"ekoa-certificate-generator/internal/utils"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

type Certificate struct {
	PK                string `json:"PK"`
	ReportId          int    `json:"reportId"`
	ContentId         int    `json:"contentId"`
	ContentSlug       string `json:"contentSlug"`
	ContentTitle      string `json:"contentTitle"`
	CourseStartedAt   string `json:"courseStartedAt,omitempty"`
	CourseFinishedAt  string `json:"courseFinishedAt,omitempty"`
	StudentId         int    `json:"studentId"`
	StudentName       string `json:"studentName"`
	StudentEmail      string `json:"studentEmail"`
	StudentGroupIds   string `json:"studentGroupIds,omitempty"`
	ExpiresAt         string `json:"expiresAt"`
	ExpirationEnabled bool   `json:"expirationEnabled"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
	FilePath          string `json:"filePath"`
}

func InterfaceToModel(data interface{}) (instance *Certificate, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)
}

func (c *Certificate) ToString() (string, error) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		log.Fatal("ERROR: parse certificate to string", err)
		return "", err
	}

	return string(jsonData), nil
}

func (c *Certificate) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                c.PK,
		"reportId":          c.ReportId,
		"contentId":         c.ContentId,
		"contentSlug":       c.ContentSlug,
		"contentTitle":      c.ContentTitle,
		"courseStartedAt":   c.CourseStartedAt,
		"courseFinishedAt":  c.CourseFinishedAt,
		"studentId":         c.StudentId,
		"studentName":       c.StudentName,
		"studentEmail":      c.StudentEmail,
		"studentGroupIds":   c.StudentGroupIds,
		"expiresAt":         c.ExpiresAt,
		"expirationEnabled": c.ExpirationEnabled,
		"createdAt":         c.CreatedAt,
		"updatedAt":         c.UpdatedAt,
		"filePath":          c.FilePath,
	}
}

func ParseDynamoAtributeToStruct(response map[string]*dynamodb.AttributeValue) (c Certificate, err error) {
	if len(response) == 0 {
		return c, errors.New("item not found")
	}
	for key, value := range response {
		if key == "id" {
			c.PK = *value.S
		}
		if key == "reportId" {
			c.ReportId, _ = strconv.Atoi(*value.N)
		}
		if key == "contentId" {
			c.ContentId, _ = strconv.Atoi(*value.N)
		}
		if key == "contentSlug" {
			c.ContentSlug = *value.S
		}
		if key == "contentTitle" {
			c.ContentTitle = *value.S
		}
		if key == "courseStartedAt" {
			c.CourseStartedAt = *value.S
		}
		if key == "courseFinishedAt" {
			c.CourseFinishedAt = *value.S
		}
		if key == "studentId" {
			c.StudentId, _ = strconv.Atoi(*value.N)
		}
		if key == "studentName" {
			c.StudentName = *value.S
		}
		if key == "studentEmail" {
			c.StudentEmail = *value.S
		}
		if key == "studentGroupIds" {
			c.StudentGroupIds = *value.S
		}
		if key == "expiresAt" {
			c.ExpiresAt = *value.S
		}
		if key == "expirationEnabled" {
			c.ExpirationEnabled = *value.BOOL
		}
		if key == "createdAt" {
			c.CreatedAt = *value.S
		}
		if key == "updatedAt" {
			c.UpdatedAt = *value.S
		}
		if key == "filePath" {
			c.FilePath = *value.S
		}
	}

	return c, nil
}

func (c *Certificate) GetFilterPK() map[string]interface{} {
	return map[string]interface{}{"PK": c.PK}
}

func (c *Certificate) GetFilterReportId() (condition expression.Expression, err error) {
	keyCond := expression.Key("reportId").Equal(expression.Value(c.ReportId))
	condition, err = expression.NewBuilder().WithKeyCondition(keyCond).Build()

	if err != nil {
		return condition, err
	}

	return condition, nil
}

func (c *Certificate) GetFilterEmail() (condition expression.Expression, err error) {
	keyCond := expression.Key("studentEmail").Equal(expression.Value(c.StudentEmail))
	condition, err = expression.NewBuilder().WithKeyCondition(keyCond).Build()

	if err != nil {
		return condition, err
	}

	return condition, nil
}

func (c *Certificate) Bytes() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Certificate) GenerateID() {
	id := uuid.NewString()
	c.PK = id
}

func (c *Certificate) SetCreatedAt() {
	c.CreatedAt = utils.GetDateTimeNowFormatted()
}

func (c *Certificate) SetUpdatedAt() {
	c.UpdatedAt = utils.GetDateTimeNowFormatted()
}

func (c *Certificate) SetFilePath() {
	c.FilePath = "pdf/" + c.StudentEmail + "/" + c.PK + ".pdf"
}

func GetTimeFormat() string {
	return "2006-01-02T15:04:05-0700"
}
