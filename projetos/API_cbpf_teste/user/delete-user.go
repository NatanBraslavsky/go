package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/mongo"
)

func DeleteByID(ctx context.Context, id bson.ObjectID) error {
	result, err := mongo.ApiProject.Collection("user").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not find")
	}

	return nil
}
