package organization

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Organization struct{}

func Routes(route *gin.Engine) {
	router := route.Group("/organization")

	router.GET("/:org_uid", func(c *gin.Context) {
		fmt.Print(c.Params)

		c.JSON(200, gin.H{
			"message": "organization",
		})
	})
}
