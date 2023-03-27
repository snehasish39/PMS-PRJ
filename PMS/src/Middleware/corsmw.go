package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// EnableCrossDomain .
func EnableCrossDomain(c *gin.Context) {
	fmt.Println("CORS")
	cors.Middleware(cors.Config{
		Origins: "*",
		Methods: "GET, PUT, POST, DELETE",
		// Methods: "GET, PUT, POST, DELETE, OPTIONS, TRACE",
		RequestHeaders: "Origin, Authorization, Content-Type, Cookies, responseType, Accept, ClientID",
		// RequestHeaders:  "Origin, Authorization, Content-Type,Content-Length, Cookies, responseType, Accept, X-CSRF-Token, Accept-Encoding, X-Header,X-Y-Header",
		MaxAge:          5000 * time.Second,
		Credentials:     true,
		ValidateHeaders: true,
		ExposedHeaders:  "Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma",
		// ExposedHeaders: "Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma, Link, X-Header,X-Y-Header",
	})
	return
}
