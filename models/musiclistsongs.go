package models

import "github.com/google/uuid"

type MusicListSongs struct {
	MusicListID uuid.UUID `json:"Id" gorm:"type:uuid;"`
	SongId      uuid.UUID `json:"SongId" gorm:"type:uuid;"`
	Score       int
	Songs       []Songs `gorm:"foreignKey:SongId;references:ID"`
}

func (mls *MusicListSongs) SaveMusicListSong() (*MusicListSongs, error) {

	var err error
	err = DB.Create(&mls).Error
	if err != nil {
		return &MusicListSongs{}, err
	}
	return mls, nil
}
