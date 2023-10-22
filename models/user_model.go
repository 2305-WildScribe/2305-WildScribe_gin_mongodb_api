package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	User_id     primitive.ObjectID `json:"user_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" validate:"required"`
	Password    string             `json:"password,omitempty" validate:"required"`
}

