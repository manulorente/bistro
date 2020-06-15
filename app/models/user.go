package models

import (
	"time"
	"errors"
	"strings"
)

type User struct {
    Userid 	    int 		`json:"userid" orm:"primary_key"`
    Email       string      `json:"email" orm:"size(128)"`
    Username 	string 		`json:"username" orm:"size(32)`
	Password 	string 		`json:"password" orm:"size(64)"`
	CreatedAt   time.Time 	`json:"created_at" orm:"auto"`
    UpdatedAt   time.Time 	`json:"updated_at" orm:"auto"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var UsersList = []User{
    {Userid: 0, Email: "user1@gmail.com", Username: "user1", Password: "pass1", CreatedAt: time.Now(), UpdatedAt: time.Time{}},
    {Userid: 1, Email: "user2@gmail.com", Username: "user2", Password: "pass2", CreatedAt: time.Now(), UpdatedAt: time.Time{}},
    {Userid: 2, Email: "user3@gmail.com", Username: "user3", Password: "pass3", CreatedAt: time.Now(), UpdatedAt: time.Time{}},
}

func IsUserValid(username, password string) bool {
    for _, u := range UsersList {
        if u.Username == username && u.Password == password {
            return true
        }
    }
    return false
}

// Register a new user with the given username and password
func RegisterNewUser(username, email, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
        return nil, errors.New("The password can't be empty")
    } else if !IsUsernameAvailable(username) {
        return nil, errors.New("The username isn't available")
    } else if !IsEmailAvailable(email) {
        return nil, errors.New("The email isn't available")
    }

    u := User{
        Userid: length(UsersList)+1,
        Email: email, 
        Username: username, 
        Password: password, 
        CreatedAt: time.Now(), 
        UpdatedAt: time.Time{},
    }

    //UsersList = append(UsersList, u)
    _, err := ORM.Insert(&u)

    if err == nil {
        return &u, nil
    }else{
        return nil, errors.New("The user could not be inserted")
    }
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

// Check if the supplied email is available
func IsEmailAvailable(email string) bool {
    for _, u := range UsersList {
        if u.Email == email {
            return false
        }
    }
    return true
}