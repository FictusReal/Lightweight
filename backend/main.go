package main

import (
	"net/http"

	//dbinitialize "github.com/FictusReal/Lightweight/backend/database/database"
	"github.com/FictusReal/Lightweight/backend/database"
	"github.com/gin-gonic/gin"
)

func main() {

	//dbinitialize.DBInit()
	database.DBInit()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:id", getUserID)
	r.POST("/login", authMe)
	r.POST("/user/create", makeUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
