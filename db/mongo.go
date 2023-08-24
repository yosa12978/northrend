package db

import (
	"context"
	"sync"

	"github.com/yosa12978/northrend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once sync.Once
	db   *mongo.Database
)

func GetDB() *mongo.Database {
	once.Do(func() {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(config.Config.Db.Uri).SetServerAPIOptions(serverAPI)
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}
		db = client.Database(config.Config.Db.DbName)
	})
	return db
}

func DisconnectDb() {
	if err := GetDB().Client().Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
