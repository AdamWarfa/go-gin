package main

import (
	"time"

	"github.com/AdamWarfa/go-gin/initializers"
	"github.com/AdamWarfa/go-gin/internal/controllers"
	"github.com/AdamWarfa/go-gin/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})

		time.AfterFunc(24*time.Hour, func() {
		controllers.SetWord()
	})
	router := gin.Default()

    router.GET("/", controllers.GetWord)
	router.POST("/users", controllers.SaveUser)

    router.Run(":6969")
}
