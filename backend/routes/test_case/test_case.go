package test_case

import "github.com/gin-gonic/gin"

type TestCase struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/test-case")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test case",
		})
	})
}
