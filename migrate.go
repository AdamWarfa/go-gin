package main

import (
	"github.com/AdamWarfa/go-gin/initializers"
	"github.com/AdamWarfa/go-gin/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func migrate() {
	initializers.DB.AutoMigrate(&models.User{})
}
