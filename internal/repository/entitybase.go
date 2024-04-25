package repository

import (
	"time"

	"github.com/google/uuid"
)

type EntityInterface interface {
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	GetMap() map[string]interface{}
	GetFilterId() map[string]interface{}
}

type EntityBase struct {
	PK        string    `json:"PK,omitempty"`
	SK        string    `json:"SK,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (b *EntityBase) GenerateID() {
	id := uuid.NewString()
	b.PK = id
	b.SK = id
}

func (b *EntityBase) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *EntityBase) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}

func GetTimeFormat() string {
	return "2006-01-02T15:04:05-0700"
}
