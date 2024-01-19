package main

import (
	"time"

	"github.com/AdamWarfa/go-gin/initializers"
	"github.com/AdamWarfa/go-gin/internal/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {

		time.AfterFunc(24*time.Hour, func() {
		controllers.SetWord()
	})
	router := gin.Default()

    router.GET("/", controllers.GetWord)

    router.Run(":6969")
}
