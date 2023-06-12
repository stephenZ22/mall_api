package controllers

import (
	"MallApi/db"
	"MallApi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type ProjectController struct{}

type CreateProjectJson struct {
	StoreID uint   `gorm:"column:store_id" json:"store_id"`
	Name    string `gorm:"column:name" json:"name"`
	Price   uint   `gorm:"type:decimal(9,2);column:price" json:"price"`
	Status  uint   `gorm:"default:0;column:status" json:"status"`
}

type ProjectInfoJson struct {
	ID        uint      `gorm:"column:id" json:"id"`
	StoreID   uint      `gorm:"column:store_id" json:"store_id"`
	Name      string    `gorm:"column:name" json:"project_name"`
	Price     uint      `gorm:"type:decimal(9,2);column:price" json:"project_price"`
	Status    uint      `gorm:"default:0;column:status" json:"project_status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	var newProject models.Project
	var projectJson CreateProjectJson

	err := c.ShouldBindJSON(&projectJson)
	newProject.Name = projectJson.Name
	newProject.Price = projectJson.Price
	newProject.StoreId = projectJson.StoreID
	newProject.Status = projectJson.Status

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Params error to create project",
			"error":   err.Error(),
		})

		return
	}

	if err := db.MainDb.Create(&newProject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massge": "DB error to create projects",
			"error":  err.Error(),
		})

		return
	}

	pp := ProjectInfoJson{}
	mapstructure.Decode(newProject, pp)
	c.JSON(http.StatusOK, pp)

}
