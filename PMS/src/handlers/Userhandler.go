package handlers

import (
	"PMS/controller"
	"PMS/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(engine *gin.Engine) {

	fmt.Println("USER settupuser router")
	//creating a new gin.router group
	r := engine.Group("support")
	//registers a GET route for getting all user
	r.GET("/", controller.GetAllUser)
	//registers a POST route for adding user
	r.POST("/User/Add", controller.AddUser)
	//registers a POST route for deleting user
	r.POST("/User/Delete", controller.DeleteUser)
	//registers a POST route for updating a usertype
	r.POST("/User/UpdateType", controller.UpdateUserByType)
	//registers a POST route for login of a user
	r.POST("/User/login", logger.Login)

}
