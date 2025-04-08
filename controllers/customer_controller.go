package controllers

import (
	"net/http"

	"main.go/schemas"
	"main.go/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerController struct {
	service services.CustomerService
}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

// @Summary Create a new customer
// @Description Create a new customer with the provided details
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer body schemas.CustomerRequest true "Customer Data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /customers [post]
func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	var req schemas.CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponseError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	customer, err := cc.service.CreateCustomer(c, req)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, "Failed to create customer")
		return
	}

	NewResponseSuccess(c, http.StatusCreated, customer)
}

// @Summary Get all customers
// @Description Retrieve a list of all customers
// @Tags Customers
// @Produce json
// @Success 200 {array} schemas.CustomerResponse
// @Router /customers [get]
func (cc *CustomerController) GetAllCustomers(c *gin.Context) {
	customers, err := cc.service.GetAllCustomers(c)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, "Failed to retrieve customers")
		return
	}

	NewResponseSuccess(c, http.StatusOK, customers)
}

// @Summary Get a customer by ID
// @Description Retrieve a customer using its ID
// @Tags Customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} schemas.CustomerResponse
// @Failure 404 {object} ErrorResponse
// @Router /customers/{id} [get]
func (cc *CustomerController) GetCustomerByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	customer, err := cc.service.GetCustomerByID(c, id)
	if err != nil {
		NewResponseError(c, http.StatusNotFound, "Customer not found")
		return
	}

	NewResponseSuccess(c, http.StatusOK, customer)
}

// @Summary Update a customer
// @Description Update customer details using ID
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param customer body schemas.CustomerRequest true "Updated Customer Data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /customers/{id} [put]
func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var req schemas.CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		NewResponseError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	if err := cc.service.UpdateCustomer(c, id, req); err != nil {
		NewResponseError(c, http.StatusInternalServerError, "Failed to update customer")
		return
	}

	NewResponseSuccess(c, http.StatusOK, "Customer updated successfully")
}

// @Summary Delete a customer
// @Description Remove a customer using its ID
// @Tags Customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /customers/{id} [delete]
func (cc *CustomerController) DeleteCustomer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if err := cc.service.DeleteCustomer(c, id); err != nil {
		NewResponseError(c, http.StatusInternalServerError, "Failed to delete customer")
		return
	}

	NewResponseSuccess(c, http.StatusOK, "Customer deleted successfully")
}
