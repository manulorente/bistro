package controllers

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

