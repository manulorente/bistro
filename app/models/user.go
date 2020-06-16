package models

import (
	"time"
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    UserID 	    int 	    `json:"userid" gorm:"primary_key"`
    Email       string      `json:"email" validate:"min=4,max=32"`
    Username 	string 		`json:"username" validate:"min=4,max=32,regexp=^[a-zA-Z]*$"`
	Password 	string 		`json:"password" validate:"min=4,max=32"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
    Products    []Product   `json:"products,omitempty"`
}

/*
func IsUserValid(username, password string) bool {
    UsersList := GetAllUsers()
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
}*/