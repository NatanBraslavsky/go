package user

import (
	"context"

	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/mongo"
)

func Insert(ctx context.Context, newUser models.User) error {
	_, err := mongo.ApiProject.Collection("user").InsertOne(ctx, newUser)

	return err
}
