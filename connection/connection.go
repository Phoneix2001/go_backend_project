package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
 Client *mongo.Client
 ToDoDB *mongo.Database
 UsersDB *mongo.Database
)

func Mongodb() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	clientDB, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	Client = clientDB
	ToDoDB = Client.Database("TodoDB")
    UsersDB = clientDB.Database("users")
	command := bson.D{{Key: "create", Value: "todos"}}

	var result bson.M

	if err := ToDoDB.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		log.Fatal(err)
	}

	collectionNames, err := ToDoDB.ListCollectionNames(context.TODO(), struct{}{})

	if err != nil {
		log.Fatalf("Failed to list collections: %v", err)
	}

	fmt.Printf("Collections in database %s: %v\n", ToDoDB.Name(), collectionNames)

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
}
