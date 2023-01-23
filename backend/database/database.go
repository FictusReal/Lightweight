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

func AddFriend(id string, friendID string, client *mongo.Client) int {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	friendObjID, fErr := primitive.ObjectIDFromHex(friendID)
	if fErr != nil {
		panic(fErr)
	}
	var user User
	var friend User
	coll := client.Database("Lightweight").Collection("users")
	err = coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&user)
	if err != nil {
		return 404
	}
	err = coll.FindOne(context.TODO(), bson.D{{"_id", friendObjID}}).Decode(&friend)
	if err != nil {
		return 404
	}
	user.Friends = append(user.Friends, friendID)
	friend.Friends = append(friend.Friends, id)
	// remove(user.Friends, friendId)

	update := bson.D{{"$set", bson.D{{"friends", user.Friends}}}}
	result, err := coll.UpdateByID(context.TODO(), objID, update)
	if err != nil {
		panic(err)
	}
	update = bson.D{{"$set", bson.D{{"friends", friend.Friends}}}}
	result, err = coll.UpdateByID(context.TODO(), friendObjID, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return 200
}

func RemoveFriend(id string, friendID string, client *mongo.Client) int {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	friendObjID, fErr := primitive.ObjectIDFromHex(friendID)
	if fErr != nil {
		panic(fErr)
	}
	var user User
	var friend User
	coll := client.Database("Lightweight").Collection("users")
	err = coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&user)
	if err != nil {
		return 404
	}
	err = coll.FindOne(context.TODO(), bson.D{{"_id", friendObjID}}).Decode(&friend)
	if err != nil {
		return 404
	}

	// user.Friends = append(user.Friends, friendID)
	// friend.Friends = append(friend.Friends, id)
	removeFriendFromList(&user.Friends, friendID)
	removeFriendFromList(&friend.Friends, id)

	update := bson.D{{"$set", bson.D{{"friends", user.Friends}}}}
	result, err := coll.UpdateByID(context.TODO(), objID, update)
	if err != nil {
		panic(err)
	}
	update = bson.D{{"$set", bson.D{{"friends", friend.Friends}}}}
	result, err = coll.UpdateByID(context.TODO(), friendObjID, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return 200
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

func GetAllUsers(client *mongo.Client) []User {
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

func ReadUserFromDatabase(UID string, name string, client *mongo.Client) (Response User, err error) {

	var filter primitive.D
	if UID == "" && name == "" { //empty strings no need for searching
		return User{}, nil
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

	//returnVal, _ := json.Marshal(result) No need
	returnVal := result
	fmt.Println(returnVal)

	return returnVal, nil

}

func removeFriendFromList(myFriendsList *[]string, ID string) {
	var indexFound int = -1
	var returnSlice []string
	myFriends := *myFriendsList
	for i := 0; i < len(myFriends); i++ {
		if ID == myFriends[i] {
			indexFound = i
		}
	}
	// if indexFound > 0 && len(myFriends) <= 1 {
	// 	returnSlice = []string{}
	// }
	if indexFound > 0 {
		returnSlice = append(myFriends[:indexFound], myFriends[indexFound+1:]...)
	}
	if returnSlice == nil || len(returnSlice) <= 0 {
		returnSlice = []string{}
	}

	*myFriendsList = returnSlice
}

// func getExercise(id string, client *mongo.Client) Exercise {
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(objID)
// 	coll := client.Database("Lightweight").Collection("exercises")
// 	var yeeter User
// 	err = coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&yeeter)
// 	fmt.Println(yeeter)
// 	return yeeter
// }
