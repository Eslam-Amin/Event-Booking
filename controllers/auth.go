package controllers

import (
	"net/http"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func Singup(context *gin.Context){
	var user models.User
	err := context.BindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data.",
			"error": err.Error(),
		})
		return
	}
	
	err = user.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't save user, try again later!",
			"error": err.Error(),
		})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": user,
	})
}