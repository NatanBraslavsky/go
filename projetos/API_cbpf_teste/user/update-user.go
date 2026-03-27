package user

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/mongo"
)

func UpdateByID(ctx context.Context, id bson.ObjectID, updatedUser models.User) error {
	result, err := mongo.ApiProject.Collection("user").UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updatedUser},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}

	return nil
}
