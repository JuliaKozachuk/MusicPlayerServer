package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB(postgres string) {

	db, err := gorm.Open("postgres", postgres)
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&MusicList{})
	db.AutoMigrate(&MusicListSongs{})
	db.AutoMigrate(&Songs{})

	DB = db

}
