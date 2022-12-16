package main

import (
	"context"
	"fmt"
	mongorm "github.com/aliforever/go-mongolio"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var err error

	var client *mongo.Client
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return
	}

	db := client.Database("mydb")

	Users = mongorm.Collection[User](db, "users")
	Accounts = mongorm.Collection[Account](db, "accounts")

	user, err := Users.FindOne(bson.M{"first_name": "Ali"})
	if err != nil {
		fmt.Println(err)
		return
	}

	account, err := Accounts.FindOne(bson.M{"username": "admin"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", user)
	fmt.Printf("%#v\n", account)
}
