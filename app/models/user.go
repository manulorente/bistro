package models

import (
	"time"
	"errors"
	"strings"
)

type User struct {
    Username 	string 		`json:"username"`
	Password 	string 		`json:"-"`
	CreatedAt   time.Time 	`json:"created_at"`
	UpdatedAt   time.Time 	`json:"updated_at"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var UsersList = []User{
    {Username: "user1", Password: "pass1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
    {Username: "user2", Password: "pass2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
    {Username: "user3", Password: "pass3", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// Register a new user with the given username and password
func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
        return nil, errors.New("The password can't be empty")
    } else if !IsUsernameAvailable(username) {
        return nil, errors.New("The username isn't available")
    }

    u := User{Username: username, Password: password, CreatedAt: time.Now(), UpdatedAt: time.Now()}

    UsersList = append(UsersList, u)

    return &u, nil
}

// Check if the supplied username is available
func IsUsernameAvailable(username string) bool {
    for _, u := range UsersList {
        if u.Username == username {
            return false
        }
    }
    return true
}

func IsUserValid(username, password string) bool {
    for _, u := range UsersList {
        if u.Username == username && u.Password == password {
            return true
        }
    }
    return false
}