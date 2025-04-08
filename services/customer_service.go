package services

import (
	"context"
	"main.go/models"
	"main.go/repo"
	"main.go/schemas"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, req schemas.CustomerRequest) (schemas.CustomerResponse, error)
	GetCustomerByID(ctx context.Context, id primitive.ObjectID) (schemas.CustomerResponse, error)
	UpdateCustomer(ctx context.Context, id primitive.ObjectID, req schemas.CustomerRequest) error
	DeleteCustomer(ctx context.Context, id primitive.ObjectID) error
	GetAllCustomers(ctx context.Context) ([]schemas.CustomerResponse, error)
}

type customerService struct {
	repo repo.CustomerRepository
}

func NewCustomerService(repo repo.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) CreateCustomer(ctx context.Context, req schemas.CustomerRequest) (schemas.CustomerResponse, error) {
	customer := models.Customer{
		FullName:   req.FullName,
		FullAdress: req.FullAdress,
		Phone:      req.Phone,
		Email:      req.Email,
	}
	id, err := s.repo.CreateCustomer(ctx, customer)
	if err != nil {
		return schemas.CustomerResponse{}, err
	}
	customer.ID = id
	return schemas.CustomerResponse(customer), nil
}

func (s *customerService) DeleteCustomer(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.DeleteCustomer(ctx, id)
}

func (s *customerService) GetAllCustomers(ctx context.Context) ([]schemas.CustomerResponse, error) {
	customers, err := s.repo.GetAllCustomers(ctx)
	if err != nil {
		return nil, err
	}

	var customerResponses []schemas.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, schemas.CustomerResponse{
			ID:         customer.ID,
			FullName:   customer.FullName,
			FullAdress: customer.FullAdress,
			Phone:      customer.Phone,
			Email:      customer.Email,
		})
	}

	return customerResponses, nil
}

func (s *customerService) GetCustomerByID(ctx context.Context, id primitive.ObjectID) (schemas.CustomerResponse, error) {
	customer, err := s.repo.GetCustomerByID(ctx, id)
	if err != nil {
		return schemas.CustomerResponse{}, err
	}

	return schemas.CustomerResponse{
		ID:         customer.ID,
		FullName:   customer.FullName,
		FullAdress: customer.FullAdress,
		Phone:      customer.Phone,
		Email:      customer.Email,
	}, nil
}

func (s *customerService) UpdateCustomer(ctx context.Context, id primitive.ObjectID, req schemas.CustomerRequest) error {
	customer := models.Customer{
		FullName:   req.FullName,
		FullAdress: req.FullAdress,
		Phone:      req.Phone,
		Email:      req.Email,
	}

	return s.repo.UpdateCustomer(ctx, id, customer)
}