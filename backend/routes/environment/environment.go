package environment

import "github.com/gin-gonic/gin"

type Environment struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/environment")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "environment",
		})
	})
}
