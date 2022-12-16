package main

import "github.com/aliforever/go-mongolio"

var Accounts *mongolio.C[Account]

type Account struct {
	ID       int    `bson:"_id,omitempty"`
	UserID   int    `bson:"user_id"`
	Username string `bson:"username"`
}
