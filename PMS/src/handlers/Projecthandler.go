package handlers

import (
	controller "PMS/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupProjectRouter(engine *gin.Engine) {

	fmt.Println("SETUPPROJECTROUTER")
	r := engine.Group("support")
	//registers a GET route for getting all projects
	r.GET("/Project/All", controller.GetAllProject)
	//registers a POST route for adding a project
	r.POST("/Project/Addproject", controller.AddProject)
	//registers a POST route for deleting a project
	r.POST("/Project/delete", controller.DeleteProject)
	//register a POST route for updating a projectstatus
	r.POST("/Project/Update", controller.UpdateProject)
	//registers a POST route for assigning a project
	r.POST("/Project/assign", controller.AssignProject)

}
