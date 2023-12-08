package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dirFiles = "./files"
	file = "users-file.json"
	path = dirFiles + "/" + file
	uriDb = "mongodb://localhost:27017"
	db = "go-concepts"
	collection = "users-service"
	qtdReadToDb = 3
)

var isReading bool = false

type User struct {
	Id        int `bson:"id,omitempty" json:"id"`
	Uuid      string `bson:"uid,omitempty" json:"uid"`
	FirstName string `bson:"first_name,omitempty" json:"first_name"`
	LastName  string `bson:"last_name,omitempty" json:"last_name"`
	Username  string `bson:"username,omitempty" json:"username"`
	Email     string `bson:"email,omitempty" json:"email"`
	Phone     string `bson:"phone_number,omitempty" json:"phone_number"`
	Birthday  string `bson:"date_of_birth,omitempty" json:"date_of_birth"`
	Password  string `bson:"password,omitempty" json:"password"`
}

func dbConnect() (mClient *mongo.Client, mCollection *mongo.Collection, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriDb))
	if err != nil { return }

	mCollection = client.Database(db).Collection(collection)
	return
}

func addUsersToDb(mCollection *mongo.Collection, users []User) (res *mongo.InsertManyResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	iUsers := make([]interface{}, len(users))
	for i := range users { iUsers[i] = users[i] }

	res, err = mCollection.InsertMany(ctx, iUsers)
	if err != nil { return }

	return
}

func getUsersFromApi() (users []User) {
	rand.Seed(time.Now().UnixNano())

	qtdUsers := rand.Intn(15) + 1
	apiUser := fmt.Sprintf("https://random-data-api.com/api/v2/users?response_type=json&size=%d", qtdUsers)

	res, err := http.Get(apiUser)
	if err != nil { panic(err) }

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil { panic(err) }

	json.Unmarshal([]byte(body), &users)
	return
}

func getUsersFromFile() (users []User) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { return []User{} }
		panic(err)
	}

	json.Unmarshal([]byte(f), &users)
	return 
}

func resetFile() {
	emptyData, err := json.Marshal([]User{})
	if err != nil { panic(err) }

	err = os.WriteFile(path, emptyData, 0644)
	if err != nil { panic(err) }
}

func serviceWriteUsers(c chan <- bool, cEnd chan <- bool) {
	for {
		if(isReading) { 
			fmt.Printf("- (file | write) Waiting for read...\n")
			time.Sleep(time.Second * 3)
			continue 
		}

		usersFile := getUsersFromFile()
		usersApi := getUsersFromApi()

		users := append(usersFile, usersApi...)

		data, err := json.Marshal(users)
		if err != nil { panic(err) }

		err = os.WriteFile(path, data, 0644)
		if err != nil { panic(err) }

		fmt.Printf("- (file | write) More %v users added to file!\n", len(usersApi))
		fmt.Printf("- (file | log) Total users: %v\n", len(users))

		c <- true
		time.Sleep(time.Second * 5)
	}
}

func serviceReadUsers(c <- chan bool, cEnd chan <- bool) {
	countRead := 0

	_, collection, err := dbConnect()
	if err != nil { 
		fmt.Printf("- (db) %v\n", err)
		cEnd <- true

		return
	}

	for range c {
		countRead++
		if(!(countRead % qtdReadToDb == 0)) { continue }

		isReading = true
		countRead = 0

		users := getUsersFromFile()
		if len(users) == 0 { 
			fmt.Printf("- (file | read) No users in file!\n")
			isReading = false
			continue 
		}

		res, err := addUsersToDb(collection, users)
		if err != nil {
			fmt.Printf("- (db) Users not added to the database: %v\n", err)
			isReading = false
			continue
		}

		fmt.Printf("- (db) Inserted %v document(s)!\n", len(res.InsertedIDs))
		fmt.Printf("- (db) Total users: %v\n", len(users))

		resetFile()
		fmt.Printf("- (file | read) File cleaned!\n")

		isReading = false
	}	
}

func main() {
	c := make(chan bool)
	cEnd := make(chan bool)

	go serviceWriteUsers(c, cEnd)
	go serviceReadUsers(c, cEnd)

	<- cEnd
}