package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/controllers"
	"main.go/repo"
	"main.go/services"
)

func CustomerRoutes(router *gin.Engine) {
	customerGroup := router.Group("/customers")
	
	mongo:= &mongo.Collection{}
	repo:= repo.NewCustomerRepo(mongo)
	services:= services.NewCustomerService(repo)
	controllers:= controllers.NewCustomerController(services)

	{
		customerGroup.POST("/", controllers.CreateCustomer)
		customerGroup.GET("/", controllers.GetAllCustomers)
		customerGroup.GET("/:id", controllers.GetCustomerByID)
		customerGroup.PUT("/:id", controllers.UpdateCustomer)
		customerGroup.DELETE("/:id", controllers.DeleteCustomer)
	}
}
