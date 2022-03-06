package main

import mongorm "github.com/aliforever/go-mongorm2"

var Accounts *mongorm.C[Account]

type Account struct {
	ID       int
	UserID   int
	Username string
}

func (Account) Name() string {
	return "accounts"
}
