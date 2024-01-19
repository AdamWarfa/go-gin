package main

import (
	"time"

	"github.com/AdamWarfa/go-gin/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

		time.AfterFunc(24*time.Hour, func() {
		controllers.SetWord()
	})
	router := gin.Default()

    router.GET("/", controllers.GetWord)

    router.Run(":6969")
}
