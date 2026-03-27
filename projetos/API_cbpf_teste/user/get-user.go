package user

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/mongo"
)

func GetByID(ctx context.Context, id bson.ObjectID) (mongo.User, error) {
	var foundUser mongo.User

	err := mongo.ApiProject.Collection("user").FindOne(ctx, bson.M{"_id": id}).Decode(&foundUser)
	if err != nil {
		return mongo.User{}, err
	}

	return foundUser, nil
}
