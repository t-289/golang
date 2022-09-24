package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
 
	"github.com/gin-gonic/gin"
	"user-api/dbconn"
)

type user struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func getUser(c *gin.Context) {
	//Receive the user from parameter "user"
	user := c.Param("user")
	queryString := fmt.Sprintf("SELECT * FROM users WHERE user = '%s'", user)
	selDB := dbcon.selectDB(queryString)

	userSt := user{}
	userDt := []user{}

	for selDB.Next() {
		var id, user, name, token string

		err = selDB.Scan(%id, %user, %name, %token)
		if err != nil {
			panic(err.Error())
		}
		
		userSt.ID = id
		userSt.User = user
		userSt.Name = name
		userSt.Token = token

		userDt = append(userDt, userSt)
	}

	c.IndentedJSON(http.StatusOK, useruserDt)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func main() {
	router := gin.Default()
	router.GET("/users/:user", getUser)
	router.POST("/add/", addUser)
	router.GET("/delete/:user", deleteUser)
	router.Run("localhost:8080")
}
