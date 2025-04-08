package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerResponse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName   string             `json:"full_name" bson:"full_name"`
	FullAdress string             `json:"full_address" bson:"full_address"`
	Phone      string             `json:"phone" bson:"phone"`
	Email      string             `json:"email" bson:"email"`
}
