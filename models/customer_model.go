package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FullName   string             `bson:"name"`
	FullAdress string             `bson:"full_address"`
	Phone      string             `bson:"phone"`
	Email      string             `bson:"email"`
}
