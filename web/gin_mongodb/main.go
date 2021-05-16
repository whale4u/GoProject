package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//do some operation
	// 指定获取要操作的数据集
	collection := client.Database("xuxiaoxian").Collection("user")
	user1 := User{
		UserName: "Niko",
		Password: "Niko",
	}
	user2 := User{
		UserName: "Niko1",
		Password: "Niko1",
	}
	user3 := User{
		UserName: "Niko2",
		Password: "Niko2",
	}
	// insert user1
	insertResult, err := collection.InsertOne(context.TODO(), user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	// insert user2, user3
	users := []interface{}{user2, user3}
	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//close connection
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

//https://www.liwenzhou.com/posts/Go/go_mongodb/
