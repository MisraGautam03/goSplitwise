package main

import (
	"github.com/gin-gonic/gin"
)

var bank = make(map[int]map[int]float32)
var mapID = make(map[int]Person)
var maxID = 1

func main() {
	router := gin.New()
	router.POST("/register", registerUser)
	router.POST("/recordTransaction", recordTransaction)
	router.GET("/balance/:id", getBalance)
	err := router.Run()
	if err != nil {
		return
	}
}
