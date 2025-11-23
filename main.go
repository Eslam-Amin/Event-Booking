package main

import (
	"net/http"

	"example.com/event-booking/db"
	"example.com/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	
	server.GET("/ping", func (context *gin.Context){
	
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
