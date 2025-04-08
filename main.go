package main

import (
	"main.go/conn"
	"main.go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoDB := conn.NewMongoDB("mongodb://localhost:27017", "easy_os")
	
	r := gin.Default()
	routes.Handler(r, mongoDB)

	r.Run()
}
