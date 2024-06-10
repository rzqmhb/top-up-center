package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rzqmhb/top-up-center/service"
)

type ItemAPI interface {
	GetAll(c *gin.Context)
	GetByGame(c *gin.Context)
	GetByKeywords(c *gin.Context)
}

type itemAPI struct {
	itemService service.ItemService
}

func NewItemAPI(itemService service.ItemService) ItemAPI {
	return &itemAPI{itemService: itemService}
}

func (i *itemAPI) GetAll(c *gin.Context) {
	items, err := i.itemService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (i *itemAPI) GetByGame(c *gin.Context) {
	pathParam := c.Param("gameId")
	gameId, err := strconv.Atoi(pathParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"invalid game id"})
		return
	}

	items, err := i.itemService.GetByGame(gameId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (i *itemAPI) GetByKeywords(c *gin.Context) {
	keywords := c.Query("keyword")
	items, err := i.itemService.GetByKeywords(keywords)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}