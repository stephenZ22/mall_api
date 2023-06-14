package controllers

import (
	"MallApi/db"
	"MallApi/logger"
	"MallApi/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreController struct{}

type StoreError struct {
	code    int
	message string
}

type NewStore struct {
	UserId   uint `json:"user_id"`
	Status   uint `json:"status"`
	StoreNum uint `json:"store_num"`
}

type UpdateStore struct {
	Status uint `json:"status"`
}

func (sc *StoreController) CreateStore(c *gin.Context) {
	var newStore NewStore

	if err := c.ShouldBindJSON(&newStore); err != nil {
		logger.Error("Create Store Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create Store Error",
			"error":   err.Error(),
		})
		return
	}

	s := models.Store{}
	s.UserId = newStore.UserId
	s.Status = newStore.Status
	s.StoreNum = newStore.StoreNum

	if err := db.MainDb.Save(&s).Error; err != nil {
		logger.Error("Save Store Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create Store Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"result":  newStore,
	})
}

func (sc *StoreController) UpdateStore(c *gin.Context) {
	us := UpdateStore{}
	if err := c.ShouldBindJSON(&us); err != nil {
		logger.Error("Update Store params error", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update Store params error",
			"error":   err.Error(),
		})

		return
	}

	if err := checkStoreStatus(us); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "status value error",
			"error":   err.Error(),
		})

		return
	}

	s := models.Store{}

	db.MainDb.First(&s, c.Param("id"))
	s.Status = us.Status
	db.MainDb.Save(&s)
	result, err := convertStoreRes(s)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "convert json error",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"result":  result,
	})
}

func convertStoreRes(s models.Store) (NewStore, error) {
	var sr NewStore

	sr.Status = s.Status
	sr.UserId = s.UserId
	sr.StoreNum = s.StoreNum

	return sr, nil
}

func checkStoreStatus(us UpdateStore) error {
	fmt.Println("US Status is %d\n", us.Status)
	if us.Status > 1 {
		err := &StoreError{message: "status value error", code: 0}
		return err
	}

	return nil
}

func (e *StoreError) Error() string {
	return e.message
}
