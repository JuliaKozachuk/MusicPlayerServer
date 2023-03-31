package models

import "github.com/google/uuid"

type MusicList struct {
	Id             uuid.UUID        `json:"Id" gorm:"type:uuid;default:gen_random_uuid()",gorm:"primary_key"`
	Name           string           `json:"Name" gorm:"not null;unique" json:"name"`
	UserID         uuid.UUID        `json:"UserID" gorm:"type:uuid;"`
	MusicListSongs []MusicListSongs `gorm:"foreignKey:MusicListID;references:Id"`
}

func (ml *MusicList) SaveMusicList() (*MusicList, error) {

	var err error
	err = DB.Create(&ml).Error
	if err != nil {
		return &MusicList{}, err
	}
	return ml, nil
}
