package route

import (
	"misic_play/controllers"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func SetupRouter(sess *session.Session) {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

	router.POST("/upload", controllers.UploadImage)

	_ = router.Run(":9888")
}
