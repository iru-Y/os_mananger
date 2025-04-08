package main

import (
	"github.com/gin-gonic/gin"
	"main.go/routes"
)

func main() {
	r := gin.Default()

	routes.CustomerRoutes(r)

	r.Run(":8080")
}
