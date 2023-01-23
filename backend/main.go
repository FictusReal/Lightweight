package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// "reflect"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/FictusReal/Lightweight/backend/database"
	"github.com/gin-gonic/gin"
)

func main() {

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
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := database.DBInit()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("DEFEK")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user/create", func(c *gin.Context) {
		type LoginReq struct {
			Name  string `json:"name"`
			Token string `json:"token"`
		}
		var req LoginReq
		c.BindJSON(&req)
		user := database.CreateUser(req.Name, []string{}, []string{}, client)
		c.JSON(200, user)
	})

	//GET /user/:id
	r.GET("/user/:id", func(c *gin.Context) {
		var name database.User
		id := c.Param("id")
		fmt.Println(id)
		name = database.GetUser(id, client)
		// json, err := database.ReadUserFromDatabase("", name.Name, client)
		if err != nil {
			panic(err)
		}
		c.JSON(200, name)
	})

	r.GET("/user", func(c *gin.Context) {
		allofthem := database.GetAllUsers(client)

		c.JSON(200, allofthem)
	})

	r.GET("/user/:id/workout", func(c *gin.Context) {
		var name database.User
		id := c.Param("id")
		fmt.Println(id)
		name = database.GetUser(id, client)
		// json, err := database.ReadUserFromDatabase("", name.Name, client)
		if err != nil {
			panic(err)
		}
		c.JSON(200, name.Workouts)
	})

	r.GET("/user/:id/friend", func(c *gin.Context) {
		var name database.User
		id := c.Param("id")
		fmt.Println(id)
		name = database.GetUser(id, client)
		// json, err := database.ReadUserFromDatabase("", name.Name, client)
		if err != nil {
			panic(err)
		}
		c.JSON(200, name.Friends)
	})

	//GET /user/:id
	r.GET("/user/:id/friend/:friendID", func(c *gin.Context) {
		var name database.User
		friendID := c.Param("friendID")
		name = database.GetUser(friendID, client)
		// json, err := database.ReadUserFromDatabase("", name.Name, client)
		if err != nil {
			panic(err)
		}
		c.JSON(200, name)
	})

	r.PUT("/user/:id/friend/:friendID", func(c *gin.Context) {
		id := c.Param("id")
		friendID := c.Param("friendID")
		statusCode := database.AddFriend(id, friendID, client)
		c.JSON(statusCode, "")
	})

	r.DELETE("/user/:id/friend/:friendID", func(c *gin.Context) {
		id := c.Param("id")
		friendID := c.Param("friendID")
		statusCode := database.RemoveFriend(id, friendID, client)
		c.JSON(statusCode, "")
	})

	//POST /user/:id/workout   -   Create a workout
	// r.POST("/user/:id/workout",func(c *gin.Context){
	// 	var req []string
	// 	c.BindJSON(&req)
	// 	var name database.User
	// 	id := c.Param("id")
	// 	fmt.Println(id)
	// 	name = database.GetUser(id, client)
	// 	// json, err := database.ReadUserFromDatabase("", name.Name, client)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	c.JSON(200, user)
	// })

	// r.GET("/user/:id", getUserID)
	// r.POST("/login", authMe)
	// r.POST("/user/create", makeUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
