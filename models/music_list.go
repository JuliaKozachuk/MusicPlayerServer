package models

import "github.com/google/uuid"

type MusicList struct {
	Id   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()",gorm:"primary_key`
	Name string    `gorm: not null; unique" json:"name"`

	UserID int ``
}

func (ml *MusicList) SaveMusicList() (*MusicList, error) {

	var err error
	err = DB.Create(&ml).Error
	if err != nil {
		return &MusicList{}, err
	}
	return ml, nil
}
