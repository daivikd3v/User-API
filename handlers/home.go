package handlers

import (
	"github.com/gin-gonic/gin"
)

//Home Handler Struct
type Home struct {
}

//GetHomeHandler creates an instance of Home handler Struct.
func GetHomeHandler() *Home {
	return &Home{}
}

//Index - returns a welcome message in JSON format.
func (home Home) index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You have reached the home page route",
	})
}

//Get - returns a message indicating a get route has reached.
func (home Home) get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You have reached the get route",
	})
}
