// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	dirFiles string = "./files"
// 	file string = "creating-file.txt"
// 	path string = dirFiles + "/" + file
// )

// type User struct {
// 	Id        string `bson:"id,omitempty" json:"id"`
// 	Uuid      string `bson:"uid,omitempty" json:"uid"`
// 	FirstName string `bson:"first_name,omitempty" json:"first_name"`
// 	LastName  string `bson:"last_name,omitempty" json:"last_name"`
// 	Username  string `bson:"username,omitempty" json:"username"`
// 	Email     string `bson:"email,omitempty" json:"email"`
// 	Phone     string `bson:"phone_number,omitempty" json:"phone_number"`
// 	Birthday  string `bson:"date_of_birth,omitempty" json:"date_of_birth"`
// 	Password  string `bson:"password,omitempty" json:"password"`
// }

// func dbConnect() (mClient *mongo.Client, mCollection *mongo.Collection, err error) {
// 	uriDb := "mongodb://localhost:27017"
// 	db := "go-concepts"
// 	collection := "users-service"

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriDb))
// 	if err != nil {
// 		return
// 	}

// 	mCollection = client.Database(db).Collection(collection)
// 	return
// }

// func addUsersToDb(mCollection *mongo.Collection, users []User) (res *mongo.InsertManyResult, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()

// 	iUsers := make([]interface{}, len(users))
// 	for i := range users {
// 		iUsers[i] = users[i]
// 	}

// 	res, err = mCollection.InsertMany(ctx, iUsers)
// 	if err != nil {
// 		return
// 	}

// 	// fmt.Printf("Inserted %v document(s)!\n", len(res.InsertedIDs))
// 	return
// }

// func main() {
// 	fmt.Printf("Hello, world!\n")
// }