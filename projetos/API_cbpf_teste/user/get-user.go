package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	mongodriver "go.mongodb.org/mongo-driver/v2/mongo"

	"exemplo.com/api-cbpf/models"
	mongoconn "exemplo.com/api-cbpf/mongo"
)

func GetByID(ctx context.Context, id bson.ObjectID) (models.User, error) {
	var foundUser models.User

	err := mongoconn.ApiProject.Collection("user").FindOne(ctx, bson.M{"_id": id}).Decode(&foundUser)
	if err != nil {
		if errors.Is(err, mongodriver.ErrNoDocuments) {
			return models.User{}, ErrUserNotFound
		}

		return models.User{}, err
	}

	return foundUser, nil
}
