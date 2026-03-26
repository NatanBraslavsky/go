package env

import "os"

type EnvConfig struct {
	MongoUrl string
}

//nolint:gochecknoglobals
var C *EnvConfig

func Load() {
	C = &EnvConfig{
		MongoUrl: os.Getenv("MONGO_URL"),
	}
}
