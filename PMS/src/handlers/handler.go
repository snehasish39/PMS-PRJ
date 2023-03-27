package handlers

import (
	middleware "PMS/Middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {

	fmt.Println("SetupRoutes")

	//two lines register two middleware functions to the Gin engine
	engine.Use(middleware.DBConnectionMiddleware)
	engine.Use(middleware.AuthMiddleware)
	// engine.Use(middleware.EnableCrossDomain)

	// engine.Use(cors.Default())
	//functions set up the routes for handling user and project-related requests
	SetupUserRouter(engine)
	SetupProjectRouter(engine)

}
