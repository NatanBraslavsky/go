package env

import "os"

type EnvConfig struct {
	MongoUrl string
}

var C *EnvConfig

func Load() {
	C = &EnvConfig{
		MongoUrl: os.Getenv("MONGO_URL"),
	}
}
