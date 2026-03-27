package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/mongo"
)

func Insert(ctx context.Context, newUser mongo.User) error {
	count, err := mongo.ApiProject.Collection("user").CountDocuments(ctx, bson.M{"email": newUser.Email})
	if err != nil {
		return errors.New("erro genérico")
	}

	if count > 1 {
		return errors.New("email já cadastrado")
	}

	_, err = mongo.ApiProject.Collection("user").InsertOne(ctx, newUser)
	if err != nil {
		return errors.New("erro ao inserir o usuário")
	}

	return nil
}
