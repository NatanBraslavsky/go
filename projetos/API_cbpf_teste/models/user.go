package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID    bson.ObjectID 	`json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             	`json:"name" binding:"required"`
	Email string             	`json:"email" binding:"required"`
	Password string          	`json:"password" binding:"required"`
}

