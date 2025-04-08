package repo

import (
	"context"
	"main.go/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer models.Customer) (primitive.ObjectID, error)
	GetCustomerByID(ctx context.Context, id primitive.ObjectID) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, id primitive.ObjectID, customer models.Customer) error
	DeleteCustomer(ctx context.Context, id primitive.ObjectID) error
	GetAllCustomers(ctx context.Context) ([]models.Customer, error)
}
