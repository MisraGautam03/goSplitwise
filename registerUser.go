package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func registerUser(c *gin.Context){
	var person Person
	err := c.BindJSON(&person)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters",
		})
		return
	}
	person.ID = maxID
	mapID[maxID] = person
	maxID++
	c.JSON(http.StatusOK,person)
}