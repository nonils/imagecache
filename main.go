package main

import (
	"agileengine/imagecache/pkg/utils"
	"github.com/joho/godotenv"
)
import "github.com/gin-gonic/gin"

func main() {
	godotenv.Load()
	port := utils.GetConfigValueFromKey(utils.Port)
	utils.LoadImages()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + port)
}
