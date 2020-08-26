package data

import (
	"encoding/json"
	"errors"
	"time"

	guuid "github.com/google/uuid"
)

//Users map stores users in Memory
var users map[guuid.UUID]*User

//User model
type User struct {
	Uuid       guuid.UUID `json:"uuid"`
	Name       string     `json:"name" validate:"required,min=3,max=25"`
	Age        int        `json:"age" validate:"required,min=0,max=125"`
	Percentage float32    `json:"percentage" validate:"required,min=0,max=100`
	Time       time.Time  `json:"time" validate:"required"`
}

func init() {
	users = make(map[guuid.UUID]*User)
}

//UnmarshalJSON processes time before User model is encoded by Default json decoder.
func (u *User) UnmarshalJSON(data []byte) (err error) {
	var basicUser struct {
		Name       string
		Age        int
		Percentage float32
		Time       string
	}

	if err = json.Unmarshal(data, &basicUser); err != nil {
		return err
	}

	u.Name = basicUser.Name
	u.Age = basicUser.Age
	u.Percentage = basicUser.Percentage
	time, err := time.Parse("2006-01-02 15:04:05", basicUser.Time)
	u.Time = time
	return err
}

//MarshalJSON processes time before User model is encoded by Default json encoder.
func (u User) MarshalJSON() ([]byte, error) {
	basicUser := struct {
		UUID       guuid.UUID
		Name       string
		Age        int
		Percentage float32
		Time       string
	}{
		UUID:       u.Uuid,
		Name:       u.Name,
		Age:        u.Age,
		Percentage: u.Percentage,
		Time:       u.Time.Format("2006-01-02 15:04:05"),
	}

	return json.Marshal(basicUser)
}

//Create creates a new user and adds it to the user map.
func (u *User) Create() {
	u.Uuid = guuid.New()
	users[u.Uuid] = u
}

//Update updates an already existing user in the map.
func (u *User) Update() error {

	if _, ok := users[u.Uuid]; !ok {
		return errors.New("user with UUID doesn't exist")
	}

	users[u.Uuid] = u
	return nil
}

//Delete deletes the User from the map.
func (u *User) Delete() error {
	var user *User
	user, ok := users[u.Uuid]
	if !ok {
		return errors.New("user with UUID doesn't exist")
	}
	u.Name = user.Name
	u.Age = user.Age
	u.Percentage = user.Percentage
	u.Time = user.Time
	delete(users, u.Uuid)
	return nil
}
