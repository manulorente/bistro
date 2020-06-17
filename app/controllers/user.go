package controllers

import (
	"math/rand"
    "net/http"
    "strconv"
    //"time"
	// Local
    "github.com/manulorente/bistro/models"
    
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

func ShowLoginPage(c *gin.Context) {
	Render(c, gin.H{
        "title": "Login",
    }, "login.html")
}

func (server *Server) Register(c *gin.Context) {
    var err error
    // Create new user and validate it
    user := models.User{}
    // Obtain the posted username and password values
    user.Username = c.PostForm("username")
    user.Email = c.PostForm("email")
    user.Password = c.PostForm("password")
	user.Prepare("register")
    err = user.Validate("register")

    if err == nil {
        userCreated, err := user.SaveUser(server.DB)

        if err == nil {
            // If the user is created, set the token in a cookie and log the user in
            token := GenerateSessionToken()
            c.SetCookie("token", token, 3600, "", "", false, true) 
            c.Set("is_logged_in", true)

            c.HTML(200, "menu.html", gin.H{
            "MsgTitle":   "Registration Successful",
            "MsgContent": "Welcome " + userCreated.Username})
        } else {
            // If the username/password combination is invalid, show the error message on the login page
            c.HTML(http.StatusBadRequest, "register.html", gin.H{
                "ErrorTitle":   "Registration Failed",
                "ErrorMessage": err.Error()})
        }
    }else{
        // If the username/password combination is invalid, show the error message on the login page
        c.HTML(http.StatusBadRequest, "register.html", gin.H{
            "ErrorTitle":   "Registration Failed",
            "ErrorMessage": err.Error()})        
    }
}

func (server *Server) Login(c *gin.Context) {
    var err error
    user := models.User{}
    user.Username = c.PostForm("username")
    user.Password = c.PostForm("password")
	user.Prepare("login")
    err = user.Validate("login")

    if err == nil {
        userOK, err := user.CheckUser(server.DB)
    
        if err == nil {
            token := GenerateSessionToken()
            //c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
            c.SetCookie("token", token, 3600, "", "", false, true)

            c.Set("is_logged_in", true)
            c.Redirect(http.StatusTemporaryRedirect, "/:userOK.id/view")
        } else {
            c.HTML(http.StatusBadRequest, "login.html", gin.H{
                "ErrorTitle":   "Login Failed",
                "ErrorMessage": "User not found"})
        }
    } else {
        c.HTML(http.StatusBadRequest, "login.html", gin.H{
            "ErrorTitle":   "Login Failed",
            "ErrorMessage": "Invalid credentials provided"})
    }
}

func Logout(c *gin.Context) {

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
    c.Redirect(http.StatusTemporaryRedirect, "/")
}

