package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t-289/golang/user-api/dbconn"
)

type User struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func getUser(c *gin.Context) {
	//Receive the user from parameter "user"
	user := c.Param("user")
	queryString := fmt.Sprintf("SELECT * FROM users WHERE token = '%s'", user)
	selDB, err := dbconn.DBSelect(queryString)

	if err != nil {
		panic(err.Error())
	}

	userSt := User{}
	userDt := []User{}

	for selDB.Next() {
		var id, users, name, token string

		err = selDB.Scan(&id, &users, &name, &token)
		if err != nil {
			panic(err.Error())
		}

		userSt.ID = id
		userSt.User = user
		userSt.Name = name
		userSt.Token = token

		userDt = append(userDt, userSt)
	}

	if len(userDt) != 0 {
		c.IndentedJSON(http.StatusOK, userDt)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func delUser(c *gin.Context) {
	user := c.Param("user")
	queryString := fmt.Sprintf("DELETE FROM users WHERE token = '%s'", user)
	result, err := dbconn.DBDelete(queryString)

	if err != nil {
		panic(err.Error())
	}

	if result > 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"Status": "Deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})

}

func main() {
	router := gin.Default()
	router.GET("/users/:user", getUser)
	router.GET("/delete/:user", delUser)
	router.Run("localhost:8080")
}
