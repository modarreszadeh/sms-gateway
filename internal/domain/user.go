package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	FullName    string             `bson:"fullName"`
	PhoneNumber string             `bson:"phoneNumber"`
}

func NewUser(fullName, phoneNumber string) *User {
	return &User{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}
