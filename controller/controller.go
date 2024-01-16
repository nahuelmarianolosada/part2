package controller

import (
	"net/http"
	"part2/repository"
	"part2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var numberService service.NumberService

func init() {
	repo := &repository.NumberCollection{}
	numberService = &service.NumberServiceImpl{repo}
}

func GetAll(c *gin.Context) {
	list := numberService.GetAll()

	// Check if the list is empty
	if list == nil || len(list) == 0 {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, list)
}

func GetType(c *gin.Context) {
	number := c.Param("number")
	numberStr, err := strconv.Atoi(number)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	numberType := numberService.GetByID(numberStr)

	c.JSON(http.StatusOK, numberType)
}

func Insert(c *gin.Context) {
	var body NumberRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := numberService.Insert(body.Number)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "")
}

type NumberRequest struct {
	Number int `json:"number"`
}
