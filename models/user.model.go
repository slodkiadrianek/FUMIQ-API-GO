package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"firstName" bson:"firstName"`
	LastName    string             `json:"lastName" bson:"lastName"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	IsActivated bool               `json:"isActivated" bson:"isActivated"`
}
