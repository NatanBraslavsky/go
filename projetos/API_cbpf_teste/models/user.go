package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string        `binding:"required"   json:"name"`
	Email    string        `binding:"required"   json:"email"`
	Password string        `binding:"required"   json:"password"`
}
