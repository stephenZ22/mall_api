package controllers

import (
	"MallApi/db"
	"MallApi/logger"
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
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if err := db.MainDb.Where("email =?", loginData.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "email not found ",
			})

			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user.Password)
	fmt.Println(loginData.Password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email

	exTime := time.Now().Add(30 * time.Minute)
	claims["exp"] = exTime.Unix()

	secretKey := os.Getenv("SecretKey")
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken, "message": "ok"})
}

func (uc *UserController) GetAllUser(c *gin.Context) {
	var users []models.User

	err := db.MainDb.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    users,
	})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	fmt.Println(string(hash))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bcrypt Password"})
	}

	newUser := models.User{Email: body.Email, Password: string(hash)}

	if err := db.MainDb.Omit("UserInfo").Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		return
	}
	fmt.Println(newUser.Password)
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
