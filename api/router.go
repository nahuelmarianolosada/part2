package api

import (
	"net/http"
	"part2/controller"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/", controller.Insert)

	r.GET("/", controller.GetAll)
	r.GET("/:number", controller.GetType)

	r.Run(":8082")
}
