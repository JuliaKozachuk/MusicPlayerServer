package main

import (
	"fmt"
	"misic_play/connectaws"
	"misic_play/connectdb"
	"misic_play/controllers"

	"github.com/gin-gonic/gin"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string
var MyBucket string
var filepath string
var filename string
var key string

func main() {

	connectaws.LoadEnv()
	// LoadEnv()
	connectdb.InitRouter()

	awsAccessKeyID := connectaws.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	fmt.Println("My access key ID is ", awsAccessKeyID)

	sess := connectaws.ConnectAws()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()

	})

	router.POST("/upload", controllers.UploadImage)

	router.GET("/download/:item", controllers.DownloadFromS3Bucket)

	_ = router.Run(":9888")

}
