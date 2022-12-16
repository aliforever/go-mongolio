package main

import (
	mongorm "github.com/aliforever/go-mongolio"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Users *mongorm.C[User]

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
}

func (u *User) FindAccounts() (acs []Account, err error) {
	acs, err = Accounts.Find(bson.M{
		"user_id": u.ID,
	})
	return
}
