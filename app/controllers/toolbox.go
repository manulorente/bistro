package controllers

import (
	"regexp"
	"errors"
	"strings"
	"runtime"
	"github.com/gin-gonic/gin"
	"net/http"

)
var Regex_var = regexp.MustCompile("^/(register|login|view|create)/([a-zA-Z0-9]+)$")

func GetFnName() string {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	return extractFnName.ReplaceAllString(functionObject.Name(), "$1")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)
	
	switch c.Request.Header.Get("Accept") {
	case "application/json":
	  // Respond with JSON
	  c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
	  // Respond with XML
	  c.XML(http.StatusOK, data["payload"])
	default:
	  // Respond with HTML
	  c.HTML(http.StatusOK, templateName, data)
	}
  
}

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("User name Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
