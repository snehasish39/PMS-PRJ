package controller

import (
	"PMS/Models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Projects struct {
	ProjectId     int    `json:"projectID";pk`
	ProjectTitle  string `json:"projectTitle"`
	AssigneeId    int    `json:"assigneeID"`
	ProjectStatus string `json:"projectStatus"`
}

func GetAllProject(c *gin.Context) {
	fmt.Println("INSIDEGETALLPROJECT")
	const functionName = "controller.getAllProject"

	p, err := Models.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal sever error")
		return
	}

	fmt.Println("DONE")
	c.JSON(http.StatusOK, p)
}

func AddProject(c *gin.Context) {
	fmt.Println("ADD PROJECT")
	const functionName = "controller.AddProject"
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	fmt.Println("STEP1")

	var dumi Projects
	err = json.Unmarshal(body, &dumi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to unmarshal ")
		return
	}
	fmt.Println("STEP2")

	var mProject Models.YpProjectData
	status := "active"
	mProject = Models.YpProjectData{
		ProjectId:     dumi.ProjectId,
		ProjectTitle:  dumi.ProjectTitle,
		AssigneeId:    dumi.AssigneeId,
		ProjectStatus: status,
	}
	fmt.Println(mProject)
	id, err := Models.AddProject(&mProject)
	fmt.Println("FINALSTEP")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to add")
		return
	}
	fmt.Println("DONE")

	c.JSON(http.StatusOK, gin.H{"Added project": id})

}

func DeleteProject(c *gin.Context) {
	fmt.Println("inside the delete project")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	var u Models.YpProjectData
	err = json.Unmarshal(body, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
	}
	fmt.Println("u data", u)
	fmt.Println(u.ProjectId)

	projectdata, err := Models.GetProjectdata(u.ProjectId)
	if err != nil {
		fmt.Println("ERR in getprojectdata", err)
	}

	fmt.Println("project data:", projectdata)
	fmt.Printf(projectdata.ProjectStatus)
	if projectdata.ProjectStatus == "active" || projectdata.ProjectStatus == "inDev" || projectdata.ProjectStatus == "inProgress" || projectdata.ProjectStatus == "inTesting" {
		fmt.Println("inside yp")

		projectdata.ProjectStatus = "inactive"

		err = Models.UpdateProject(projectdata)
		if err != nil {
			fmt.Println("print err", err)
			c.JSON(http.StatusInternalServerError, "failed to delete")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})

	}
	c.JSON(http.StatusOK, u)

}

func UpdateProject(c *gin.Context) {

	fmt.Println("inside update user")
	const functionName = "controller.UpdateProject"
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	var u Models.YpProjectData
	err = json.Unmarshal(body, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to unmarshal")
	}
	fmt.Println("PRINT THE PROJECT DATA", u)
	fmt.Println("ProjectData", u.ProjectId)
	fmt.Println("ProjectStatus", u.ProjectStatus)

	projectdata, err := Models.GetProjectdata(u.ProjectId)
	if err != nil {
		fmt.Println("error in get user", err)
	}
	fmt.Println("print project data", projectdata)
	fmt.Println(projectdata.ProjectStatus)

	if projectdata.ProjectStatus == "active" || projectdata.ProjectStatus == "inDev" || projectdata.ProjectStatus == "inTesting" || projectdata.ProjectStatus == "inProgress" {
		fmt.Println("inside yp")

		projectdata.ProjectStatus = u.ProjectStatus

		err = Models.UpdateProject(projectdata)
		if err != nil {
			fmt.Println("print err", err)
			c.JSON(http.StatusInternalServerError, "failed to update")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Project status updated successfully"})

	}
	c.JSON(http.StatusOK, u)
	fmt.Println("Updated project")

}

func AssignProject(c *gin.Context) {
	fmt.Println("AssignPRJ")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	var myData map[string]int
	err = json.Unmarshal(body, &myData)
	fmt.Println("SHOW BODYVALUE", myData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "coould not find the project")
		return
	}

	uid, ok := myData["userId"]
	fmt.Println("U DATA", uid)

	if !ok {
		c.JSON(http.StatusInternalServerError, "faild to get user data")
		return
	}

	pid, ok := myData["projectId"]
	fmt.Println("SHOW PID", pid)

	if !ok {
		c.JSON(http.StatusInternalServerError, "faild to get project data")
		return
	}

	project := Models.YpProjectData{ProjectId: pid}
	fmt.Println("SHOW PROJECT", project)
	err = Models.AssignProject(uid, project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "faild to assign data to project")
		return
	}

	// auprj := Models.YpUserData{UserId: uid}
	// fmt.Println("PROJECT", auprj)
	// err = Models.AssignUserProject(pid, auprj)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, "failed to assign data to user")
	// 	return
	// }

	c.JSON(http.StatusOK, "Updated!!")
}
