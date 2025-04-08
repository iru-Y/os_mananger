package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/conn"
	"main.go/controllers"
	_ "main.go/docs"
	"main.go/repo"
	"main.go/services"
)

func Handler(router *gin.Engine, mongoDB *conn.MongoDB) {
	baseRoute := "/api/v1"

	customerGroup := router.Group(fmt.Sprintf("%s/customers", baseRoute))

	repository := repo.NewCustomerRepo(mongoDB.DB.Collection("customers"))
	service := services.NewCustomerService(repository)
	controller := controllers.NewCustomerController(service)

	{
		customerGroup.POST("/", controller.CreateCustomer)
		customerGroup.GET("/", controller.GetAllCustomers)
		customerGroup.GET("/:id", controller.GetCustomerByID)
		customerGroup.PUT("/:id", controller.UpdateCustomer)
		customerGroup.DELETE("/:id", controller.DeleteCustomer)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
