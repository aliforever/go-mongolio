package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/aliforever/go-mongolio"
)

func main() {
	var err error

	var client *mongo.Client
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return
	}

	db := client.Database("mydb")

	Users = mongolio.Collection[User](db, "users")
	Accounts = mongolio.Collection[Account](db, "accounts")

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
