package database

import (
	"context"
	"encoding/json"
	"fmt"

	//"os"

	//"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type Restaurant struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	RestaurantId string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       interface{}
}

func TestReadDatabase(client *mongo.Client) {

	coll := client.Database("sample_restaurants").Collection("restaurants")
	filter := bson.D{{"cuisine", "Hamburgers"}}

	var result Restaurant
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	// end findOne

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}

// func()
func CreateUser(name string, friends []string, workouts []string, client *mongo.Client) User {
	coll := client.Database("Lightweight").Collection("users")
	newUser := struct {
		Name     string
		Friends  []string
		Workouts []string
	}{
		Name:     name,
		Friends:  friends,
		Workouts: workouts,
	}
	fmt.Println("FEKKK")
	result, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}

	newId := result.InsertedID
	id := newId.(primitive.ObjectID)

	userStruct := User{ID: id, Name: name, Friends: friends, Workouts: workouts}
	fmt.Println(userStruct)
	return userStruct
}

func GetUser(id string, client *mongo.Client) User {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(objID)
	coll := client.Database("Lightweight").Collection("users")
	var yeeter User
	err = coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&yeeter)
	fmt.Println(yeeter)
	return yeeter
}

func GetAllUsers(client *mongo.Client) []User{
	coll := client.Database("Lightweight").Collection("users")
	var results []User
	filter := bson.D{{}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
		}
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func ReadUserFromDatabase(UID string, name string, client *mongo.Client) (Response []byte, err error) {

	var filter primitive.D
	if UID == "" && name == "" { //empty strings no need for searching
		return nil, nil
	}
	coll := client.Database("Lightweight").Collection("users")
	if UID != "" && name == "" {
		filter = bson.D{{"ID", UID}}
		fmt.Println("UID")
	} else if UID == "" && name != "" {
		fmt.Println("NAMEEEEE")
		filter = bson.D{{"name", name}}
	} else {
		filter = bson.D{{"name", name}, {"ID", UID}}
	}

	var result User
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
			//return
		}
		panic(err)
	}

	returnVal, _ := json.Marshal(result)
	fmt.Println(string(returnVal))

	return returnVal, nil

}
