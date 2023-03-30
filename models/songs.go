package models

import "github.com/google/uuid"

type Songs struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()",gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Autor       string    `json:"autor"`
	Link        string    `json:"link"`
}
