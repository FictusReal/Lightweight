package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Mgroup  string `json:"mgroup"`
	Ex_type string `json:"ex_type"`
}

type Workout struct {
	ID   primitive.ObjectID `bson:"_id"`
	Date primitive.DateTime
}

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Friends  []string `json:"friends"`
	Workouts []string `json:"workouts"`
}
