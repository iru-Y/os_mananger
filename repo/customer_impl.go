package repo

import (
	"context"

	"main.go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepoImpl struct {
	Collection *mongo.Collection
}

func NewCustomerRepo(collection *mongo.Collection) CustomerRepository {
	return &CustomerRepoImpl{Collection: collection}
}

func (r *CustomerRepoImpl) CreateCustomer(ctx context.Context, customer models.Customer) (primitive.ObjectID, error) {
	res, err := r.Collection.InsertOne(ctx, customer)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *CustomerRepoImpl) GetCustomerByID(ctx context.Context, id primitive.ObjectID) (*models.Customer, error) {
	var customer models.Customer
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepoImpl) UpdateCustomer(ctx context.Context, id primitive.ObjectID, customer models.Customer) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": customer})
	return err
}

func (r *CustomerRepoImpl) DeleteCustomer(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *CustomerRepoImpl) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var customer models.Customer
		if err := cursor.Decode(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
