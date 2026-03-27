package user

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/mongo"
)

func GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User

	cursor, err := mongo.ApiProject.Collection("user").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var currentUser models.User

		err = cursor.Decode(&currentUser)
		if err != nil {
			return nil, err
		}

		users = append(users, currentUser)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
