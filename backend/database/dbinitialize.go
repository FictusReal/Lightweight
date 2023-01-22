package database

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInit() {

	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://FictusReal:wqjuGiZEEjYiO1p0@cluster0.qtyok13.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RESULT!!!!!!")
	fmt.Println(reflect.TypeOf(databases))
	fmt.Println(databases)
	fmt.Println(reflect.TypeOf(databases[2]))
	fmt.Println(databases[3])
	fmt.Println(databases[4])
	fmt.Println(databases[1])

	//TestReadDatabase(client)
	// var friends []string
	//WriteUserToDatabase("Larry Wheels", friends, client)
	println("In Database now")
	//ReadUserFromDatabase("", "Larry Wheels", client)
	// CreateUser("Ronnie Coleman", []string{}, []string{}, client)

}
