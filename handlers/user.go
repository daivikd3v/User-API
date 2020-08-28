package handlers

import (
	"net/http"

	"github.com/daivikd3v/User-API/data"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
)

//User handler Struct
type User struct {
}

//GetUserHandler creates and returns an  instance of User Handler
func GetUserHandler() *User {
	return &User{}
}

//Post creates a users in memory from the request.
func (user User) post(c *gin.Context) {

	u, err := unmarshal(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	err = validate(u)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	u.Create()

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data":    u,
		})
}

//Put updates an already existing user in memory from the request.
func (user User) put(c *gin.Context) {

	u, err := unmarshal(c)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	err = validate(u)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	u.Uuid, err = guuid.Parse(c.Param("uuid"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   "Invalid UUID",
			})
		return
	}

	err = u.Update()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data":    u,
		})
}

//Delete deletes a user from UUID in the request
func (user User) delete(c *gin.Context) {

	u := data.User{}

	Uuid, err := guuid.Parse(c.Param("uuid"))

	u.Uuid = Uuid

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   "Invalid UUID",
			})
		return
	}

	err = u.Delete()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data":    u,
		})
}

// Decode json in user struct
func unmarshal(c *gin.Context) (*data.User, error) {
	var u data.User

	err := c.ShouldBindJSON(&u)

	return &u, err
}

//Validate request parameters and return error
func validate(u *data.User) error {
	validate := validator.New()
	err := validate.Struct(*u)

	return err
}
