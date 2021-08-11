package Dbconnet

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var database *mongo.Client
var datactx context.Context

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://root:gc742899@cluster0.mqkff.mongodb.net/TESTMVC?authSource=admin&replicaSet=atlas-fe50vv-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx) ///
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database, datactx = client, ctx
	// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(databases)
}

func GetDatabase() *mongo.Client {
	return database
}
func GetDatabasectx() context.Context {
	return datactx
}

// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://root:gc742899@cluster0.mqkff.mongodb.net/TEST?authSource=admin&replicaSet=atlas-fe50vv-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(databases)
