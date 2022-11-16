package test_suites

import "github.com/gin-gonic/gin"

type TestSuites struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/test-suites")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test-suites",
		})
	})
}
