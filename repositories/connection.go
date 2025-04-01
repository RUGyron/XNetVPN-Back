package repositories

import (
	"XNetVPN-Back/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"sync"
	"time"
)

var (
	once           sync.Once
	MajorityClient *mongo.Client
	SimpleClient   *mongo.Client
)

// ConnectToMongoDB Singleton MongoClient
func ConnectToMongoDB() {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.TimeoutMongoConnect)*time.Millisecond)
		defer cancel()
		journal := true

		majorityOptions := options.Client().ApplyURI(config.Config.MongoUri)
		majorityOptions.SetWriteConcern(&writeconcern.WriteConcern{W: "majority", Journal: &journal})
		majorityOptions.SetReadConcern(readconcern.Majority())
		majorityOptions.SetReadPreference(readpref.Primary())
		majorityOptions.SetMinPoolSize(5)
		majorityOptions.SetMaxPoolSize(10)

		simpleOptions := options.Client().ApplyURI(config.Config.MongoUri)
		simpleOptions.SetWriteConcern(writeconcern.W1())
		simpleOptions.SetReadConcern(readconcern.Local())
		simpleOptions.SetReadPreference(readpref.PrimaryPreferred())
		simpleOptions.SetMinPoolSize(5)
		simpleOptions.SetMaxPoolSize(10)

		var err error

		MajorityClient, err = mongo.Connect(ctx, majorityOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = MajorityClient.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		SimpleClient, err = mongo.Connect(ctx, simpleOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = SimpleClient.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")
	})
}
