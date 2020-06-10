package handlers

import (
	"regexp"
	"runtime"
	"github.com/gin-gonic/gin"
	"net/http"

)

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
