package controllers

import (
	"MallApi/db"
	"MallApi/logger"
	"MallApi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectController struct{}

type CreateProjectJson struct {
	StoreID uint    `gorm:"column:store_id" json:"store_id"`
	Name    string  `gorm:"column:name" json:"name"`
	Price   float32 `gorm:"type:decimal(9,2);column:price" json:"price"`
	Status  uint    `gorm:"default:0;column:status" json:"status"`
}

type UpdateProjectJson struct {
	Name   string  `gorm:"column:name" json:"name"`
	Price  float32 `gorm:"type:decimal(9,2);column:price" json:"price"`
	Status uint    `gorm:"column:status" json:"status"`
}

type ProjectInfoJson struct {
	ID        uint      `gorm:"column:id" json:"id"`
	StoreID   uint      `gorm:"column:store_id" json:"store_id"`
	Name      string    `gorm:"column:name" json:"project_name"`
	Price     float32   `gorm:"column:price" json:"project_price"`
	Status    uint      `gorm:"column:status" json:"project_status"`
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
		logger.Error("Params error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Params error to create project",
			"error":   err.Error(),
		})

		return
	}

	if err := db.MainDb.Create(&newProject).Error; err != nil {
		logger.Error("DB error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "DB error to create projects",
			"error":   err.Error(),
		})

		return
	}

	result := convertProjectRes(newProject)

	c.JSON(http.StatusOK, result)
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	p_id := c.Param("id")
	if err := db.MainDb.Delete(&models.Project{}, p_id).Error; err != nil {
		logger.Error("Project Delete Error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "project delete error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "ok",
		"project_id": p_id,
	})
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	p_id := c.Param("id")

	var p models.Project

	if err := db.MainDb.First(&p, p_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "find project Error",
			"error":   err.Error(),
		})
		return
	}

	var up UpdateProjectJson
	if err := c.ShouldBindJSON(&up); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "project params Error",
			"error":   err.Error(),
		})
		return
	}

	p.Status = up.Status
	p.Price = up.Price
	p.Name = up.Name

	if err := db.MainDb.Save(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "project save Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"result":  convertProjectRes(p),
	})
}

func (pc *ProjectController) GetAllProject(c *gin.Context) {
	var projects []models.Project

	if err := db.MainDb.Find(&projects).Error; err != nil {
		logger.Error("GetAllProject error:", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "GetAllProject",
			"error":   err.Error(),
		})

		return
	}
	result := convertProjects(projects)

	c.JSON(http.StatusOK, result)
}

func (pc *ProjectController) GetProjectByID(c *gin.Context) {
	var p models.Project
	if err := db.MainDb.First(&p, c.Param("id")).Error; err != nil {
		logger.Error("Find project error")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"result":  convertProjectRes(p),
	})
}

func convertProjects(ps []models.Project) []ProjectInfoJson {
	var pis []ProjectInfoJson
	for i := 0; i < len(ps); i++ {
		pis = append(pis, convertProjectRes(ps[i]))
	}

	return pis
}

func convertProjectRes(p models.Project) ProjectInfoJson {
	pi := ProjectInfoJson{}

	pi.ID = p.ID
	pi.StoreID = p.StoreId
	pi.Name = p.Name
	pi.Price = p.Price
	pi.Status = p.Status
	pi.CreatedAt = p.CreatedAt
	pi.UpdatedAt = p.UpdatedAt

	return pi
}
