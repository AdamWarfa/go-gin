package controllers

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AdamWarfa/go-gin/internal/services"
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