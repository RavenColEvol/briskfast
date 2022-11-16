package organization

import (
	"github.com/gin-gonic/gin"
)

type Organization struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/organization")

	router.POST("/", func(c *gin.Context) {})

	router.PUT("/:org_uid", func(c *gin.Context) {})

	router.GET("/:org_uid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "organization",
		})
	})

	router.DELETE("/:org_uid", func(c *gin.Context) {})
}
