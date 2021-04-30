package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	DBName = "tawb-database"
)

var databaseInstance *mongo.Database

func ConnectMongoDatabase() *mongo.Database{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		fmt.Println(err)
	}
	db := client.Database(DBName)
	databaseInstance = db // storing database instance for further usage
	fmt.Println(db.Name())
	return db
}

func GetDatabaseInstance() *mongo.Database{
	return databaseInstance
}