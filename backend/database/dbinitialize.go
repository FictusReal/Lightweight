package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func DBInit() {
// 	dsn := "postgresql://azaan:G5y2e6a8ojMkDnNoPA_tLA@sea-jackal-4744.6wr.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
// 	ctx := context.Background()
// 	conn, err := pgx.Connect(ctx, dsn)
// 	defer conn.Close(context.Background())
// 	if err != nil {
// 		log.Fatal("failed to connect database", err)
// 	}
// 	//fmt.Println("WE Are NOW Initializing Databse!!!")

// 	var now time.Time
// 	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
// 	if err != nil {
// 		log.Fatal("failed to execute query", err)
// 	}

// 	fmt.Println(now)

// }

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
	fmt.Println(databases)
}
