package model

import (
	"encoding/json"
	"log"
)

type Course struct {
	ID             string  `db:"id"`
	CurseducaIds   *string `db:"idsCurseduca"`
	ValidationDays *int    `db:"validadeEmDias"`
}

type CurseducaIds struct {
	ContentId   int    `json:"content_id"`
	ContentUuid string `json:"content_uuid"`
}

func (c *Course) GetCurseducaIds() *CurseducaIds {
	if c.CurseducaIds == nil {
		return nil
	}
	var ids CurseducaIds
	err := json.Unmarshal([]byte(*c.CurseducaIds), &ids)
	if err != nil {
		log.Printf("ERROR: Failed to unmarshal CurseducaIds: %v. Raw JSON: %s", err, *c.CurseducaIds)
		return nil
	}
	return &ids
}
