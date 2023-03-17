package controllers

import (
	"fmt"
	"misic_play/connectaws"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

func DownloadFromS3Bucket(c *gin.Context) {

	bucket := connectaws.GetEnvWithKey("BUCKET_NAME")
	fmt.Println(bucket)
	item := c.Param("item")

	//payload := []byte(item)

	file, err := os.Create(item)
	fmt.Println(file, err)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	sess := c.MustGet("sess").(*session.Session)

	downloader := s3manager.NewDownloader(sess)

	buff := &aws.WriteAtBuffer{}

	numBytes, err := downloader.Download(buff,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	//fmt.Println(bucket, key)
	if err != nil {
		fmt.Println(err)
	}

	data := buff.Bytes()
	c.JSON(http.StatusOK, gin.H{
		"data":     data,
		"fileName": file.Name(),
	})

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

}
