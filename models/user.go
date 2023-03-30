package models

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	//ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type User struct {
	Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()",gorm:"primary_key"`
	Username string    `gorm: not null; unique" json:"username"`
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
