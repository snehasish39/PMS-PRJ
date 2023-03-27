package controller

import (
	"PMS/Models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId     int    `json:"userId";pk`
	UserName   string `json:"userName"`
	UserType   string `json:"userType"`
	Password   string `json:"password"`
	UserStatus string `json:"userStatus"`
}

///func to GETALLUSER
func GetAllUser(c *gin.Context) {

	fmt.Println("INSIDE GETALLUSER")
	const functionName = "controller.getAllUser"

	fmt.Println("hey i am insde")
	u, err := Models.GetAllUsers()
	fmt.Println(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return

	}

	c.JSON(http.StatusOK, u)
	fmt.Println("GETALLUSER OK")
}

//func to ADDUSER
func AddUser(c *gin.Context) {

	fmt.Println("User Added")
	const functionName = "controller.AddUser"
	// reads the HTTP request body into a byte array named body using the ioutil package
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return

	}

	var dumi User
	err = json.Unmarshal(body, &dumi)
	fmt.Println(dumi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to unmarshal")
		return
	}

	var nurser Models.YpUserData
	nurser = Models.YpUserData{
		UserName:   dumi.UserName,
		UserType:   dumi.UserType,
		UserId:     dumi.UserId,
		Password:   dumi.Password,
		UserStatus: "active",
	}
	Password, err := bcrypt.GenerateFromPassword([]byte(dumi.Password), bcrypt.DefaultCost)
	fmt.Println("HLO bcrypt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to hash password")
		return
	}
	nurser.Password = string(Password)
	fmt.Print(nurser.Password)
	fmt.Println("INSERTED pwd", Password)
	fmt.Println("nurser data", &nurser)
	fmt.Println("UserName data", nurser.UserName)
	fmt.Println("UserType data", nurser.UserType)
	fmt.Println("UserId data", nurser.UserId)

	fmt.Println("Addnewuser")
	id, err := Models.AddUser(&nurser)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to add")
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
	fmt.Println("USER ADDED SUCCESS")
}

//func to DELETEUSER
func DeleteUser(c *gin.Context) {
	fmt.Println("inside the delete user")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	var u Models.YpUserData
	err = json.Unmarshal(body, &u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to unmarshal ")
		return
	}
	fmt.Println(u)
	fmt.Println("userid", u.UserId)

	userdata, err := Models.GetUserdata(u.UserId)
	if err != nil {
		fmt.Println("error in get get userdata", err)
		return
	}
	fmt.Println("userdata found in db", userdata)

	if userdata.UserStatus == "active" {
		fmt.Println("IN DELETE")

		userdata.UserStatus = "inactive"

		err = Models.DeleteUser(userdata)
		if err != nil {
			fmt.Println("Err", err)
			c.JSON(http.StatusInternalServerError, "failed to delete")
		}

	}
	//when user is deleted,project assigned to that user becomes null

	//-------------------func to unassign the project get called--------------
	_, err = Models.UnassignProject(u.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "not updated")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})

}

func UpdateUserByType(c *gin.Context) {

	fmt.Println("inside update user")
	const functionName = "controller.UpdateUserByType"

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	var u Models.YpUserData
	err = json.Unmarshal(body, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to unmarshal")
		return
	}
	fmt.Println("u data", u)
	fmt.Println("userid", u.UserId)

	userdata, err := Models.GetUserdata(u.UserId)
	if err != nil {
		fmt.Println("error in get get userdata", err)
		return
	}
	fmt.Println("userdata found in db", userdata)

	if userdata.UserType == "user" {
		fmt.Println("IN UPDATE")

		userdata.UserType = "admin"

		err = Models.UpdateUser(userdata)
		if err != nil {
			fmt.Println("Err", err)
			c.JSON(http.StatusInternalServerError, "failed to update")
		}

		c.JSON(http.StatusOK, u)
		fmt.Println("Updated user")
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

	}

}

/////////////////// -/////////////
