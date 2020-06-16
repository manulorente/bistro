package controllers


import (
	// Third party
	"github.com/gin-gonic/gin"
	"fmt"
)

// https://www.diycode.cc/projects/gin-gonic/gin
// /menu?id=1234&table=1&token=129388
func ViewMenu(c *gin.Context) {
	id := c.Query("id")
	table := c.Query("table")
	token := c.Query("token")

	// Address to right page
	fmt.Printf("id: %s; table: %s; token: %s", id, table, token)

}
