package controllers

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AdamWarfa/go-gin/initializers"
	"github.com/AdamWarfa/go-gin/internal/services"
	"github.com/AdamWarfa/go-gin/models"
	"github.com/gin-gonic/gin"
)

var (
	randomWord     string = services.GetWord("https://random-word-api.herokuapp.com/word?number=1")
	randomWordLock sync.Mutex
)

func GetWord(c *gin.Context) {
	randomWordLock.Lock()
	defer randomWordLock.Unlock()

	fmt.Println(randomWord)

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, `["`+randomWord+`"]`)
}

func SetWord() {
	randomWordLock.Lock()
	defer randomWordLock.Unlock()

	// Update the random word
	randomWord = services.GetWord("https://random-word-api.herokuapp.com/word?number=1")

	// Reschedule the setWord function to be called again after 10 seconds
	time.AfterFunc(10*time.Hour, func() {
		SetWord()
	})
}

func SaveUser(c *gin.Context) {
	var body struct {
		Id string `json:"id"`
		Email string `json:"email"`
		Streak int `json:"streak"`
		HiScore int `json:"hiScore"`
	}
	c.Bind(&body)

	user := models.User{Id: body.Id, Email: body.Email, Streak: body.Streak, HiScore: body.HiScore}

	result:= initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {

	id:= c.Param("id")

	var body struct {
		Id string `json:"id"`
		Email string `json:"email"`
		Streak int `json:"streak"`
		HiScore int `json:"hiScore"`
	}
	c.Bind(&body)

	user := models.User{Id: body.Id, Email: body.Email, Streak: body.Streak, HiScore: body.HiScore}

	initializers.DB.Model(&user).Where("id = ?", id).Updates(&user)


	c.JSON(http.StatusOK, gin.H{"user": user})
}