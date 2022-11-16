package main

import (
	"github.com/RavenColEvol/briskfast/routes/environment"
	"github.com/RavenColEvol/briskfast/routes/organization"
	"github.com/RavenColEvol/briskfast/routes/test_case"
	"github.com/RavenColEvol/briskfast/routes/test_suites"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	organization.Routes(router)
	environment.Routes(router)
	test_case.Routes(router)
	test_suites.Routes(router)

	router.Run()
}
