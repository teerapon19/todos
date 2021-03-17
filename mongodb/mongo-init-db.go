package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/teerapon19/todos/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database
var bctx = context.Background()

func Connect() {
	var err error
	dbUrl := fmt.Sprintf("mongodb://%s:%s", env.MongoEnv.DBUrl, env.MongoEnv.DBPort)
	Client, err = mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	etx, cancel := context.WithTimeout(bctx, 10*time.Second)
	defer cancel()
	err = Client.Connect(etx)
	if err != nil {
		panic(err)
	}
	initDatabase()
	fmt.Printf("OK, MongoDB connected to %s and %s database.\n", dbUrl, env.MongoEnv.DBName)
}

func initDatabase() {
	Database = Client.Database(env.MongoEnv.DBName)
}
