package models

import (
    "time"
    "errors"
    "strings"
    "html"
    "log"
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
	"github.com/badoux/checkmail"
)

type User struct {
    ID 	        int 	    `json:"id" gorm:"primary_key;auto_increment"`
    Email       string      `json:"email" gorm:"size:100;not null;unique"`
    Username 	string 		`json:"username" gorm:"size:100;not null;unique"`
	Password 	string 		`json:"password" gorm:"size:100;not null"`
	CreatedAt   time.Time   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
    Products    []Product   `json:"products,omitempty"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare(action string) {
    switch strings.ToLower(action) {
        case "register":
            u.ID = 0
            u.Username = html.EscapeString(strings.TrimSpace(u.Username))
            u.Email = html.EscapeString(strings.TrimSpace(u.Email))
            u.Password = html.EscapeString(strings.TrimSpace(u.Password))
            u.CreatedAt = time.Now()
            u.UpdatedAt = time.Now()
        case "login":
            u.ID = 0
            u.Username = html.EscapeString(strings.TrimSpace(u.Username))
            u.Password = html.EscapeString(strings.TrimSpace(u.Password))
            u.CreatedAt = time.Now()
            u.UpdatedAt = time.Now()
    }
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
    case "update":
        if u.Username == "" {
            return errors.New("Required Nickname")
        }
        if u.Password == "" {
            return errors.New("Required Password")
        }
        if u.Email == "" {
            return errors.New("Required Email")
        }
        if err := checkmail.ValidateFormat(u.Email); err != nil {
            return errors.New("Invalid Email")
        }
        return nil

    case "login":
        if u.Password == "" {
            return errors.New("Required Password")
        }
        if u.Username == "" {
            return errors.New("Required Email or Username")
        }
        return nil

    default:
        if u.Username == "" {
            return errors.New("Required Username")
        }
        if u.Password == "" {
            return errors.New("Required Password")
        }
        if u.Email == "" {
            return errors.New("Required Email")
        }
        if err := checkmail.ValidateFormat(u.Email); err != nil {
            return errors.New("Invalid Email")
        }
        return nil
    }
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

    var err error
    if !IsUserValid(u.Username, u.Email, db){
		return &User{}, err
    }else{
        err = db.Debug().Create(&u).Error
        if err != nil {
            return &User{}, err
        }
        return u, nil
    }
}

// Check if the supplied username and emails are available
func IsUserValid(username, email string, db *gorm.DB) bool {
    users := []User{}
	_ = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
    for _, u := range users {
        if u.Username == username || u.Email == email {
            return false
        }
    }
    return true
}

// TODO: https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121
func (u *User) CheckUser(username, password string, db *gorm.DB) (*User, error) {

    users := []User{}
	_ = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
    for _, u := range users {
        if u.Username == username || u.Email == email {
            return false
        }
    }
    return true

    var err error
    if !IsUserValid(u.Username, u.Email, db){
		return &User{}, err
    }else{
        err = db.Debug().Create(&u).Error
        if err != nil {
            return &User{}, err
        }
        return u, nil
    }
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"username":  u.Username,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
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