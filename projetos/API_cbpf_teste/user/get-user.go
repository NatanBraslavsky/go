package user

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/mongo"
)

func GetByID(ctx context.Context, id bson.ObjectID) (models.User, error) {
	var foundUser models.User

	err := mongo.ApiProject.Collection("user").FindOne(ctx, bson.M{"_id": id}).Decode(&foundUser)
	if err != nil {
		return models.User{}, err
	}

	return foundUser, nil
}
