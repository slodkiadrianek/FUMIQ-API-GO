package models

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName   string             `json:"firstName" bson:"firstName"`
	LastName    string             `json:"lastName" bson:"lastName"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	isActivated bool               `json:"isActivated" bson:"isActivated"`
}
