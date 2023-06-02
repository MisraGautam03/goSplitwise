package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func getBalance(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid Parameter",
		})
		return
	}
	_, valid := mapID[id]
	if valid ==false{
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "No User with this ID",
		})
		return
	}
	if bank[id]==nil{
		bank[id] = make(map[int]float32)
	}
	c.JSON(http.StatusOK,bank[id])
	return
}