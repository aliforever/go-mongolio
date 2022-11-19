# go-mongolio
Wrappers for go mongodb driver methods using generics to remove the redundant unmarshal calls

## Install
`go get github.com/aliforever/go-mongolio`

## Usage
- Define the accounts collection:
```go
package main

import "github.com/aliforever/go-mongolio"

var Accounts *mongolio.C[Account]

type Account struct {
	ID       int
	UserID   int
	Username string
}

func (Account) CollectionName() string {
	return "accounts"
}
```
- Initialize the db and `Accounts`
```go
func main() {
    var err error
    
    var client *mongo.Client
	
    client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
        panic(err)
    }
    
    db := client.Database("mydb")
    
    Accounts = mongorm.Collection[Account](db)	
}
```
- Find an account by username:
```go
    account, err := Accounts.FindOne(bson.M{"username": "admin"})
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Printf("%#v\n", account)
```