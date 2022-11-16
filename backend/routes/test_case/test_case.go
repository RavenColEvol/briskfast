package test_case

import "github.com/gin-gonic/gin"

type TestCase struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/test-case")

	router.POST("/", func(c *gin.Context) {})

	router.PUT("/:test_case_id", func(c *gin.Context) {})

	router.GET("/:test_case_id", func(c *gin.Context) {})

	router.DELETE("/:test_case_id", func(c *gin.Context) {})
}
