package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people", postPeople)

	router.Run("localhost:8080")

}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []person{

	{Name: "Thomas Hale", Age: 19},
	{Name: "Joseph Hale", Age: 23},
}

func getPeople(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, people)

}

func postPeople(c *gin.Context) {

	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	people = append(people, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)

}
