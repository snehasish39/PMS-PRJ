package middleware

import (
	"PMS/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func DBConnectionMiddleware(c *gin.Context) {
	const functionName = "middleware.DBConnectionMiddleware"
	fmt.Println(functionName)
	config.ConnectToDatabase()
	

}
