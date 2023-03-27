package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	functionName := "Authenticate"
	fmt.Println("Inside ", functionName)
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		fmt.Println("Basic Authentication is missing ")
		c.JSON(401, nil)
		c.Abort()
		return
	}
	userName := "abc"
	pass := "123"
	if userName == username && password == pass {
		c.Next()

	} else {
		c.JSON(401, nil)
		c.Abort()
		return
	}

}
