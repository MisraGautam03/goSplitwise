package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func recordTransaction(c *gin.Context){
	var transaction Transaction
	err := c.BindJSON(&transaction)
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters",
		})
		return
	}
	if len(transaction.Shares) > 0 {
		status:= sharesTransaction(transaction)
		if status == 404{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User with some id not found",
			})
			return
		}
		if status == -1{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Parameters",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message" : "Transaction Successfully Recorded",
		})
		return
	} 
	valid := checkValidity(transaction)
	if valid == false{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User with some ID not found",
		})
		return
	}
	numUsers := len(transaction.IDs)
	splitValue := float32(transaction.Total)/float32(numUsers)
	for _,id := range transaction.IDs{
		if id == transaction.PaidBy{
			continue
		}
		if bank[transaction.PaidBy] == nil{
			bank[transaction.PaidBy] = make(map[int]float32)
		}
		bank[transaction.PaidBy][id] += splitValue
		if bank[id]==nil{
			bank[id]=make(map[int]float32)
		}
		bank[id][transaction.PaidBy] -= splitValue
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Transaction Successfully Recorded",
	})
}