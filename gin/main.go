package main

import (
	"github.com/gin-gonic/gin"
	"golang-awesome/gin/controller"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	accController := router.Group("/v1/acc")
	{
		accController.POST("/save", controller.SaveAcc)
	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
