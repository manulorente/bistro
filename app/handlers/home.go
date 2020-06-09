package handlers

import (
	// Third party
	"github.com/gin-gonic/gin"

)

func HomePage(c *gin.Context) {
	Render(c, gin.H{
		"title": "Bistro",
		"msg":   "Landing page",
		},
		"index.html")
}

//query?restaurant=X&table=Y
func QueryStrings(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=UTF-8")
	Render(c, gin.H{
		"restaurant": c.Query("restaurant"),
		"table":      c.Query("table"),
	},
	"")
}