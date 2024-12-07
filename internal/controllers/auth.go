package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"innoversepm-backend/pkg/logger"
	"net/http"
)

type AuthController struct {
	logger *logrus.Logger
}

func NewAuthController() *AuthController {
	return &AuthController{logger: logger.Logger}
}

func (a *AuthController) GetUsers(c *gin.Context) {
	a.logger.Error()
	c.JSON(http.StatusOK, gin.H{"message": "Get all users"})
}

func (a *AuthController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create a new user"})
}
