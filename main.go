package main

import (
	"fmt"
	"misic_play/connectaws"
	"misic_play/connectdb"
	"misic_play/controllers"

	"github.com/gin-gonic/gin"
)

// func GetEnvWithKey(key string) string {
// 	return os.Getenv(key)
// }
// func LoadEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 		os.Exit(1)
// 	}
// }

// var AccessKeyID string
// var SecretAccessKey string
// var MyRegion string
// var MyBucket string
// var filepath string
// var filename string
// var key string

//	func ConnectAws() *session.Session {
//		AccessKeyID = GetEnvWithKey("AWS_ACCESS_KEY_ID")
//		SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
//		MyRegion = GetEnvWithKey("AWS_REGION")
//		sess, err := session.NewSession(
//			&aws.Config{
//				Region: aws.String(MyRegion),
//				Credentials: credentials.NewStaticCredentials(
//					AccessKeyID,
//					SecretAccessKey,
//					"", // a token will be created when the session it's used.
//				),
//			})
//		if err != nil {
//			panic(err)
//		}
//		return sess
//	}
// func SetupRouter(sess *session.Session) {
// 	router := gin.Default()

// 	router.Use(func(c *gin.Context) {
// 		c.Set("sess", sess)
// 		c.Next()
// 	})
// 	//router.LoadHTMLGlob("templates/*")

// 	// router.Get("/upload", Form)
// 	router.POST("/upload", UploadImage)

// 	// router.GET("/image", controllers.DisplayImage)

//		_ = router.Run(":9888")
//	}
// func UploadImage(c *gin.Context) {
// 	sess := c.MustGet("sess").(*session.Session)
// 	uploader := s3manager.NewUploader(sess)
// 	MyBucket = GetEnvWithKey("BUCKET_NAME")
// 	file, header, err := c.Request.FormFile("photo")
// 	filename := header.Filename
// 	//fmt.Println(sess, uploader, MyBucket, file, header, err, filename)
// 	//upload to the s3 bucket
// 	up, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(MyBucket),
// 		//ACL:    aws.String("public-read"),
// 		Key:  aws.String(filename),
// 		Body: file,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error":    "Failed to upload file",
// 			"uploader": up,
// 		})
// 		fmt.Println(err)
// 		return
// 	}
// 	filepath = "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
// 	c.JSON(http.StatusOK, gin.H{
// 		"filepath": filepath,
// 	})
// }

// func DownloadFromS3Bucket(c *gin.Context) {

// 	bucket := GetEnvWithKey("BUCKET_NAME")
// 	fmt.Println(bucket)
// 	item := c.Param("item")
// 	//payload := []byte(item)

// 	file, err := os.Create(item)
// 	fmt.Println(file, err)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer file.Close()

// 	sess := c.MustGet("sess").(*session.Session)

// 	downloader := s3manager.NewDownloader(sess)

// 	buff := &aws.WriteAtBuffer{}

// 	numBytes, err := downloader.Download(buff,
// 		&s3.GetObjectInput{
// 			Bucket: aws.String(bucket),
// 			Key:    aws.String(item),
// 		})
// 	fmt.Println(bucket, key)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	data := buff.Bytes()
// 	c.JSON(http.StatusOK, gin.H{
// 		"data":     data,
// 		"fileName": file.Name(),
// 	})

// 	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

// }
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
