package logger

import (
	"PMS/Models"
	"fmt"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



func Login(c *gin.Context) {
	fmt.Println("INSIDE LOGIN")
	var input struct {
		UserName string `json:"userName" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	fmt.Println("LOGIN FAIL")
	var user Models.YpUserData

	// Retrieve the user with the given username from the database

	o := orm.NewOrm()

	err := o.QueryTable("YpUserData").Filter("UserName", input.UserName).One(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	fmt.Println("user.password", user.Password)
	fmt.Println("inputpassword", input.Password)

	if user.UserStatus == "inactive" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"authorised": false,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authorised": true,
		"userName":   user.UserName,
		"userId":     user.UserId,
		"userType":   user.UserType,
	})
}

//////////////////////////////////////////////////////////////////////
