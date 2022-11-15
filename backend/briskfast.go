package main

import (
	"github.com/RavenColEvol/briskfast/routes/organization"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	organization.Routes(router)

	router.Run()
}
