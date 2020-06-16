package controllers

import (
	"math/rand"
    "net/http"
    "strconv"
	"strings"

    "github.com/gin-gonic/gin"
)

func GenerateSessionToken() string {
    // We're using a random 16 character string as the session token
    // This is NOT a secure way of generating session tokens
    // DO NOT USE THIS IN PRODUCTION
    return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	Render(c, gin.H{
        "title": "Register"}, "register.html")
}

func Register(c *gin.Context) {
	  // Obtain the POSTed username and password values
	  username := c.PostForm("username")
	  email := c.PostForm("email")
	  password := c.PostForm("password")

	  //var sameSiteCookie http.SameSite;

	  if _, err := RegisterNewUser(username, email, password); err == nil {
		  // If the user is created, set the token in a cookie and log the user in
		  token := GenerateSessionToken()
		  //c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		  c.SetCookie("token", token, 3600, "", "", false, true) 
		  c.Set("is_logged_in", true)
  
		 // Render(c, gin.H{
			  //"title": "Successful registration & Login"}, "login.html")
			c.HTML(200, "menu.html", gin.H{
			"MsgTitle":   "Registration Successful",
			"MsgContent": "Welcome " + username})
  
	  } else {
		  // If the username/password combination is invalid,
		  // show the error message on the login page
		  c.HTML(http.StatusBadRequest, "register.html", gin.H{
			  "ErrorTitle":   "Registration Failed",
			  "ErrorMessage": err.Error()})
  
	  }
}

func ShowLoginPage(c *gin.Context) {
	Render(c, gin.H{
        "title": "Login",
    }, "login.html")
}

func PerformLogin(c *gin.Context) {
	username := c.PostForm("username")
    password := c.PostForm("password")

	//var sameSiteCookie http.SameSite;

	// Check if the username/password combination is valid
    if IsUserValid(username, password) {
        token := GenerateSessionToken()
		//c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		c.SetCookie("token", token, 3600, "", "", false, true)

		c.Set("is_logged_in", true)
		c.Redirect(http.StatusTemporaryRedirect, "/products/view")

    } else {
        c.HTML(http.StatusBadRequest, "login.html", gin.H{
            "ErrorTitle":   "Login Failed",
            "ErrorMessage": "Invalid credentials provided"})
    }
}

func Logout(c *gin.Context) {
	//var sameSiteCookie http.SameSite;

	// Clear the cookie
	//c.SetCookie("token", "", -1, "", "", sameSiteCookie, false, true)
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
    c.Redirect(http.StatusTemporaryRedirect, "/")
}


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

    UsersList = append(UsersList, u)
    //_, err := ORM.Insert(&u)

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