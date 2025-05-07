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

func NewUser(firstName string, lastName string, email string, password string,
) User {
	return User{
		ID:          primitive.NewObjectID(),
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Password:    password,
		IsActivated: false,
	}
}
