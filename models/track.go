package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	List []ListTrack `gorm:"many2many:track_list_track;"`
}
type ListTrack struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func (track *Track) Save() (*Track, error) {
	err := DB.Create(&track).Error
	if err != nil {
		return &Track{}, err
	}
	return track, nil
}
func GetAll(track *Track) {
	result := DB.Find(&track)
	fmt.Println(result)

}
