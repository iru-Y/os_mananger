package schemas

type CustomerRequest struct {
	FullName   string `json:"full_name" bson:"full_name"`
	FullAdress string `json:"full_address" bson:"full_address"`
	Phone      string `json:"phone" bson:"phone"`
	Email      string `json:"email" bson:"email"`
}
