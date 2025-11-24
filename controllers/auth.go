package controllers

import (
	"net/http"

	"example.com/event-booking/dtos"
	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Singup(context *gin.Context){
	var signupCredentials dtos.SignupCredentials

	err := context.BindJSON(&signupCredentials)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data.",
			"error": err.Error(),
		})
		return
	}

	user := models.NewUser(signupCredentials.Name, signupCredentials.Email, signupCredentials.Password)

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

func Login(context *gin.Context){
	var user *models.User
	var loginCredentials dtos.LoginCredentials
	err := context.ShouldBindJSON(&loginCredentials)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data.",
			"error": err.Error(),
		})
		return
	}
	user, err = models.GetUserByEmail(loginCredentials.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error": err.Error(),
		})
		return
	}
	err = user.ValidateCredentials(loginCredentials.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't generate token, try again later!",
			"error": err.Error(),
		})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data": gin.H{
			"token": token,
			"user":user,
		},
	})


}