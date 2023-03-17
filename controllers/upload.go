package controllers

import (
	"fmt"
	"misic_play/connectaws"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string
var MyBucket string
var filepath string
var filename string
var key string

func UploadImage(c *gin.Context) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	MyBucket := connectaws.GetEnvWithKey("BUCKET_NAME")
	file, header, err := c.Request.FormFile("photo")
	filename := header.Filename
	//fmt.Println(sess, uploader, MyBucket, file, header, err, filename)
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		//ACL:    aws.String("public-read"),
		Key:  aws.String(filename),
		Body: file,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": up,
		})
		fmt.Println(err)
		return
	}
	filepath := "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}
