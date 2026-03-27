package user

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/mongo"
)

func GetAll(ctx context.Context) ([]mongo.User, error) {
	var users []mongo.User

	cursor, err := mongo.ApiProject.Collection("user").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}()

	for cursor.Next(ctx) {
		var currentUser mongo.User

		err = cursor.Decode(&currentUser)
		if err != nil {
			return nil, err
		}

		users = append(users, currentUser)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
