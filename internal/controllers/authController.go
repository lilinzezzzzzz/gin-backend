package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all users"})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create a new user"})
}
