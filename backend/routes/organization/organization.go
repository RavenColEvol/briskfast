package organization

import "github.com/gin-gonic/gin"

type Organization struct{}

func Routes(route *gin.Engine) {
	organization := route.Group("/organization")
	organization.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "organization",
		})
	})
}
