package main
//Made by Misra Gautam Rajeev
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var bank = make(map[int]map[int]float32)
var mapID = make(map[int]Person)
var maxID = 1

func main(){
	fmt.Println("Server is starting")
	router := gin.New()
	router.POST("/register",registerUser)
	router.POST("/recordTransaction",recordTransaction)
	router.GET("/balance/:id",getBalance)
	router.Run()
}