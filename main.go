package main

import (
	"net/http"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func (context *gin.Context){
	
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/events", func (context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"data": models.GetAllEvents(),
		})
	})

	server.Run(":8080")
}