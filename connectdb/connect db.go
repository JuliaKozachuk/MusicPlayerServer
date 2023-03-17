package connectdb

import (
	"fmt"
	"log"
	"misic_play/connectaws"
	"misic_play/models"

	"github.com/joho/godotenv"
)

func InitRouter() {

	models.ConnectDB(postgresUrl())

}
func postgresUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := connectaws.GetEnvWithKey("POSTGRES_HOST")
	port := connectaws.GetEnvWithKey("POSTGRES_PORT")
	user := connectaws.GetEnvWithKey("POSTGRES_USER")
	dbname := connectaws.GetEnvWithKey("POSTGRES_DB")
	password := connectaws.GetEnvWithKey("POSTGRES_PASSWORD")
	sslmode := connectaws.GetEnvWithKey("POSTRES_SSLMODE")

	postgres := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s ", host, port, user, dbname, password, sslmode)
	fmt.Println(postgres)

	return postgres
}
