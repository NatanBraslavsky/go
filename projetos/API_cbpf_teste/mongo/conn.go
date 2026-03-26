package mongo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"exemplo.com/api-cbpf/env"
)

var (
	ApiProject *mongo.Database
	client     *mongo.Client
)

func InitDB() {
	var err error

	client, err = mongo.Connect(options.Client().ApplyURI(env.C.MongoUrl))
	if err != nil {
		panic(err)
	}

	ApiProject = client.Database("natan-estudo")
}
