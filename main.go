package main

 import (
	"github.com/gin-gonic/gin"
	"github.com/Prakhar-Shankar/patch-pal/models"
	"github.com/Prakhar-Shankar/patch-pal/controllers"
 )

 func main(){
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/bugs", controllers.CreateBug)
	router.GET("/bugs", controllers.FindBugs)
	router.GET("/bugs/:id", controllers.FindBug)
	router.PATCH("/bugs/:id", controllers.UpdateBug)
	router.DELETE("/bugs/:id", controllers.DeleteBug)

	router.Run()
 }
