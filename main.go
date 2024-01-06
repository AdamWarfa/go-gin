package main

import (
	"github.com/AdamWarfa/go-gin/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {



	router := gin.Default()

    router.GET("/", controllers.GetSlash)

    router.Run(":6969")
}
