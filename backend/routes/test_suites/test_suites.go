package test_suites

import "github.com/gin-gonic/gin"

type TestSuites struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/test-suites")

	router.POST("/", func(c *gin.Context) {})

	router.PUT("/:test_suite_id", func(c *gin.Context) {})

	router.GET("/:test_suite_id", func(c *gin.Context) {})

	router.DELETE("/:test_suite_id", func(c *gin.Context) {})
}
