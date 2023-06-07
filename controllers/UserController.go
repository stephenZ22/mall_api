package controllers

import (
	"MallApi/db"
	"MallApi/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct{}

func (uc *UserController) UserLogin(c *gin.Context) {
	var loginData struct {
		LoginName string `json:"login_name" binding:"required"`
		PassWord  string `json:"pass_word" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if err := db.MainDb.Where("login_name =?", loginData.LoginName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Invalid credentials",
			})

			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(loginData.PassWord))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login_name"] = user.LoginName

	exTime := time.Now().Add(30 * time.Minute)
	claims["exp"] = exTime.Unix()

	secretKey := os.Getenv("SecretKey")
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bcryptPasswordByte, err := models.BcryptPassword(newUser.PassWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bcrypt password"})
	}

	newUser.PassWord = bcryptPasswordByte

	if err := db.MainDb.Omit("UserInfo").Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(200, newUser)
}

func (uc *UserController) DeletedUser(c *gin.Context) {
	// 需要做权限验证
	userId := c.Param("id")
	err := db.MainDb.Delete(&models.User{}, userId).Error

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"msg":   "ok but something wrong",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (uc *UserController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello, World"})
}
