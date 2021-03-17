package env

import (
	"fmt"
	"os"
)

type mongo struct {
	DBUrl  string
	DBPort string
	DBName string
}

var MongoEnv *mongo

func Load() {
	MongoEnv = new(mongo)
	MongoEnv.initMongoDBENV()

	fmt.Println("ENV loaded!")
}

func (env *mongo) initMongoDBENV() {
	env.DBUrl = os.Getenv("MONGO_DB_URL")
	env.DBPort = os.Getenv("MONGO_DB_PORT")
	env.DBName = os.Getenv("MONGO_DB_NAME")
}
