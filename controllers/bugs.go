package controllers

import (
	"net/http"
	"github.com/Prakhar-Shankar/patch-pal/models"
	"github.com/gin-gonic/gin"
)

type CreateBugInput struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Severity string `json:"severity"`
	Status string `json:"status"`
	Assignee string `json:"assignee"`
}

func CreateBug(c *gin.Context){
	var input CreateBugInput
	if err := c.ShouldBind(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bug := models.Bug{Title: input.Title, Description: input.Description, Serverity: input.Severity, Status: input.Status, Assignee: input.Assignee }
	models.DB.Create(&bug)

	c.JSON(http.StatusOK, gin.H{"data":bug})
}

func FindBugs(c *gin.Context){
	var bugs []models.Bug
	models.DB.Find(&bugs)

	c.JSON(http.StatusOK, gin.H{"data": bugs})
}

func FindBug(c *gin.Context){
	var bug models.Bug

	if err:= models.DB.Where("id = ?", c.Param("id")).First(&bug).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bug})
}

type UpdateBugInput struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Severity string `json:"severity"`
	Status string `json:"status"`
	Assignee string `json:"assignee"`
}

func UpdateBug(c *gin.Context){
	var bug models.Bug
	if err:= models.DB.Where("id = ?", c.Param("id")).First(&bug).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	var input UpdateBugInput
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&bug).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data":bug})
}

func DeleteBug(c *gin.Context){
	var bug models.Bug

	if err := models.DB.Where("id =?", c.Param("id")).First(&bug).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bug not found"})
		return
	}

	models.DB.Delete(&bug)

	c.JSON(http.StatusOK, gin.H{"data": "Bug delete successfully!"})
}