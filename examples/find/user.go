package main

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	mongorm "github.com/aliforever/go-mongolio"
)

var Users *mongorm.C[User]

type User struct {
	ID        bson.ObjectID `bson:"_id"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
}

func (u *User) FindAccounts() (acs []Account, err error) {
	acs, err = Accounts.Find(bson.M{
		"user_id": u.ID,
	})
	return
}
