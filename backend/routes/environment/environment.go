package environment

import "github.com/gin-gonic/gin"

type Environment struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/environment")

	router.POST("/", func(c *gin.Context) {})

	router.PUT("/:env_uid", func(c *gin.Context) {})

	router.GET("/:env_uid", func(c *gin.Context) {})

	router.DELETE("/:env_uid", func(c *gin.Context) {})
}
