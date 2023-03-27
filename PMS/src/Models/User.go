package Models

import (
	"fmt"

	"github.com/beego/beego/orm"
)

type YpUserData struct {
	UserId     int    `orm:"column(userId);pk"`
	UserName   string `orm:"column(userName);"`
	UserType   string `orm:"column(userType);"`
	Password   string `orm:"column(password);"`
	UserStatus string `orm:"column(userStatus);"`
}

//METHOD func tablename
//It returns the table name in the database where the YpUserData objects will be stored
func (u *YpUserData) Tablename() string {
	return "yp_user_data"
}

//registering the YpUserData with ORM framework
func init() {
	orm.RegisterModel(new(YpUserData))
}

//func to GETUSERDATA FROM DB
func GetUserdata(userId int) (v *YpUserData, err error) {
	o := orm.NewOrm()
	v = &YpUserData{}
	if err = o.QueryTable(new(YpUserData)).Filter("userId", userId).One(v); err == nil {
		return v, nil
	} else if err == orm.ErrNoRows && err != nil {
		return v, nil
	}
	fmt.Println("error in Getuserdata: ", err)
	return nil, err
}

//func to ADDUSER TO DB
func AddUser(u *YpUserData) (id int64, err error) {

	fmt.Println("inisde add data", u)
	fmt.Println("Add User")
	o := orm.NewOrm()
	id, err = o.Insert(u)
	if err != nil {
		fmt.Println("Err in Adduser", err)
		return
	}
	return
}

//func to GETALLUSER FROM DB
func GetAllUsers(args ...string) (u []YpUserData, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(new(YpUserData)).RelatedSel().Limit(-1).All(&u)
	if err != nil {
		fmt.Println("Invalid")
		return nil, err
	}
	return u, nil
}

//func to DELETEUSER form DB
func DeleteUser(userdata *YpUserData) error {

	fmt.Println("DeleteUSER")
	o := orm.NewOrm()
	_, err := o.Update(userdata)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil

}

// func to UPDATEUSER from DB
func UpdateUser(userdata *YpUserData) error {
	fmt.Println("UPDATE USER")

	o := orm.NewOrm()

	_, err := o.Update(userdata)
	if err != nil {
		fmt.Println("Err", err)
		return err
	}
	return nil
}

// func AssignUserProject(pid int, auprj YpUserData) error {
// 	fmt.Println("INSIDE ASSIGN PROJECT")
// 	o := orm.NewOrm()
// 	if err := o.Read(&auprj); err != nil {
// 		fmt.Println("no data")
// 		return err
// 	}
// 	auprj.Assigned_prj = pid
// 	fmt.Println("PRINT PID", pid)
// 	if _, err := o.Update(&auprj); err != nil {
// 		fmt.Println("fail to update")
// 		return err
// 	}
// 	return nil
// }
