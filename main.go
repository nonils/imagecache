package main

import (
	"agileengine/imagecache/pkg/command"
	"agileengine/imagecache/pkg/handler"
	"agileengine/imagecache/pkg/repository"
	"agileengine/imagecache/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"time"
)
import "github.com/gin-gonic/gin"

func main() {
	godotenv.Load()
	utils.Cache = cache.New(5*time.Minute, 10*time.Minute)
	port := utils.GetConfigValueFromKey(utils.Port)
	repository.InitializeMongoClient()
	command.LoadImages()
	r := gin.Default()
	r.GET("/images/:id", handler.GetImageById)
	r.GET("/search/:searchTerm", handler.SearchImageByTerm)
	r.Run(":" + port)
}
