package controllers

import (
	"misic_play/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUser struct {
	Username string
}
type CreateMusicLists struct {
	Name   string    `json:"Name"`
	UserID uuid.UUID `json:"UserID"`
}
type CreateMusicListSongs struct {
	MusicListID uuid.UUID `json:"MusicListID"`

	//Link        string    `json:"link"`
}
type Song struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Autor       string    `json:"autor"`
	SongId      uuid.UUID `json:"SongId"`
}

// type CreateList struct {
// 	Name string
// }

func CreateUsers(context *gin.Context) {
	var input CreateUser
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username}
	models.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})
}
func CreateMusicList(context *gin.Context) {
	var input CreateMusicLists
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	namemusiclist := models.MusicList{Name: input.Name, UserID: input.UserID}
	models.DB.Create(&namemusiclist)

	context.JSON(http.StatusOK, gin.H{"namemusiclist": namemusiclist})
}
func CreateMusicListSong(context *gin.Context) {
	var input CreateMusicListSongs
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	SongNameInList := models.MusicListSongs{MusicListID: input.MusicListID}
	models.DB.Create(&SongNameInList)
	context.JSON(http.StatusOK, gin.H{"SongNameInList": SongNameInList})
}

func CreateSong(context *gin.Context) {
	var input Song
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Song := models.Songs{Title: input.Title, Description: input.Description, Autor: input.Autor}
	models.DB.Create(&Song)

	context.JSON(http.StatusOK, gin.H{"Song": Song})
}
