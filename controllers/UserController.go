package controllers

import (
	"MallApi/db"
	"MallApi/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	if err := db.MainDb.Create(&newUser).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(200, newUser)
}

func (uc *UserController) DeletedUser(c *gin.Context) {

}

func (uc *UserController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello, World"})
}
