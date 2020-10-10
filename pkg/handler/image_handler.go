package handler

import (
	"agileengine/imagecache/pkg/repository"
	"agileengine/imagecache/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetImageById(c *gin.Context) {
	id := c.Param("id")
	element := repository.FindImageById(id)
	if element != nil && len(element.Id) != 0 {
		c.JSON(200, element)
		return
	}
	c.JSON(404, map[string]interface{}{"message": utils.NotFoundMessage})

}

func SearchImageByTerm(c *gin.Context) {
	searchTerm := c.Param("searchTerm")
	elements := repository.SearchImageByTerm(searchTerm)
	c.JSON(200, elements)
}
