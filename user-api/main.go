package main

import (
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
	{ID: "1", User: "Eddie", Name: "Eddie Munson", Token: ""},
	{ID: "2", User: "John", Name: "John Spartan", Token: ""},
	{ID: "3", User: "Sarah", Name: "Sarah Connor", Token: ""},
}

func getUser(c *gin.Context) {
	user := c.Param("user")
	for _, usr := range users {
		if usr.User == user {
			c.IndentedJSON(http.StatusOK, usr)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func main() {
	router := gin.Default()
	router.GET("/:user", getUser)
	router.Run("localhost:8080")
}
