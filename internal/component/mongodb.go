package component

import (
	"context"
	"log"

	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Master *mongo.Database
}

func InitMongoDB() *MongoDB {
	conf := config.Get().MongoDB
	ctx := context.Background()
	var (
		mongodb = &MongoDB{}
	)

	readPreference, err := readpref.New(
		readpref.NearestMode,
		readpref.WithMaxStaleness(conf.MongoDBMaxStaleness),
	)
	if err != nil {
		return nil
	}

	db, err := mongo.Connect(context.Background(), options.Client().
		SetAppName(conf.DBAppName).
		ApplyURI(conf.MongoDBConnectionURI).
		SetMaxPoolSize(conf.MongoDBMaxConnection).
		SetMinPoolSize(conf.MongoDBMinConnection).
		SetConnectTimeout(conf.MongoDBTimeout).
		SetSocketTimeout(conf.MongoDBTimeout).
		SetMaxConnIdleTime(conf.MongoDBMaxConnIdleTime).
		SetReadPreference(readPreference),
	)

	if err != nil {
		log.Fatalf("failed creating connection MongoDB")
	}

	if err = db.Ping(ctx, readPreference); err != nil {
		log.Fatalf("failed pinging to MongoDB")
	}

	mongodb.Master = db.Database(conf.MongoDBName)

	return mongodb

}
