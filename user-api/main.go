package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

var users = []user{
	{ID: "666", User: "Eddie", Name: "Eddie Munson", Token: "RWRkaWUtNjY2"},
	{ID: "117", User: "John", Name: "John Spartan", Token: "Sm9obi0xMTc="},
	{ID: "101", User: "Sarah", Name: "Sarah Connor", Token: "U2FyYWgtMTAx"},
}

func getUser(c *gin.Context) {
	user := c.Param("user")
	for x, usr := range users {
		if usr.User == user {
			fmt.Println(x, usr)
			c.IndentedJSON(http.StatusOK, usr)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func addUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func deleteUser(c *gin.Context) {

	user := c.Param("user")
	for x, usr := range users {
		if usr.User == user {
			fmt.Println(x, usr)
			users = append(users[:x], users[x:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func main() {
	router := gin.Default()
	router.GET("/users/:user", getUser)
	router.POST("/add/", addUser)
	router.GET("/delete/:user", deleteUser)
	router.Run("localhost:8080")
}
