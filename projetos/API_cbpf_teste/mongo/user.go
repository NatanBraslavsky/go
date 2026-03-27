package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string        `binding:"required"   json:"name"`
	Email    string        `binding:"required"   json:"email"`
	Password string        `binding:"required"   json:"password"`
}

func setupUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userColl := ApiProject.Collection("user")

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetName("email_unique"),
	}

	_, err := userColl.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(err)
	}
}
