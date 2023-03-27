package Models

import (
	"fmt"

	"github.com/beego/beego/orm"
)

type YpProjectData struct {
	ProjectId     int    `orm:"column(projectId);pk"`
	ProjectTitle  string `orm:"column(projectTitle)"`
	AssigneeId    int    `orm:"column(assigneeId)" `
	ProjectStatus string `orm:"column(projectStatus)"`
}

//register the DB in orm Framework
func init() {
	orm.RegisterModel(new(YpProjectData))
}

//METHOD func Tablename
//It returns the table name in the database where the YpProjectData objects will be stored
func (p *YpProjectData) tablename() string {
	return "yp_project_data"
}

//func to ADDPROJECT in DB
func AddProject(p *YpProjectData) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(p)
	return
}

//func to GETPROJECTDATA from DB
func GetProjectdata(projectId int) (v *YpProjectData, err error) {
	fmt.Println("GET PROJECT DATA")
	o := orm.NewOrm()
	v = &YpProjectData{}
	if err = o.QueryTable(new(YpProjectData)).Filter("projectId", projectId).One(v); err == nil {
		return v, nil
	} else if err == orm.ErrNoRows && err != nil {
		return v, nil
	}
	fmt.Println("error in Getuserdata: ", err)
	return v, err
}

//func to GETALLPROJECT from DB
func GetAllProjects(args ...string) (p []YpProjectData, err error) {

	fmt.Println("getallprj")
	o := orm.NewOrm()
	_, err = o.QueryTable(new(YpProjectData)).RelatedSel().Limit(-1).All(&p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

//func to DELETEPROJECT from DB
func DeleteProject(projectdata *YpProjectData) error {
	fmt.Println("another deltete project")
	fmt.Println("UPDATE USER")

	o := orm.NewOrm()

	_, err := o.Update(projectdata)
	if err != nil {
		fmt.Println("Err", err)
		return err
	}
	return nil

}

//func to UPDATEPROJECT from DB
func UpdateProject(projectdata *YpProjectData) error {
	fmt.Println("UPDATE USER")

	o := orm.NewOrm()

	_, err := o.Update(projectdata)
	if err != nil {
		fmt.Println("Err", err)
		return err
	}
	return nil
}

//func to UNASSIGN project to a USER
func UnassignProject(userId int) (p []YpProjectData, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(new(YpProjectData)).Filter("assigneeId", userId).RelatedSel().Limit(-1).All(&p)
	if err != nil {
		return nil, err
	}

	for _, project := range p {

		if project.ProjectStatus != "resolved" {

			// assigning 0 means we unassigning the project
			if project.ProjectStatus == "active" {
				// if project is active only then we will assign the user id
				project.AssigneeId = 0
				if _, err := o.Update(&project); err != nil {
					return p, err
				}
			}
		}

	}
	return nil, nil

}

//func to ASSIGN a project to USER
func AssignProject(uid int, project YpProjectData) error {
	fmt.Println("INSIDE ASSIGN PROJECT")
	o := orm.NewOrm()
	if err := o.Read(&project); err != nil {
		fmt.Println("no data")
		return err
	}
	project.AssigneeId = uid
	fmt.Println("PRINT UID", uid)
	if _, err := o.Update(&project); err != nil {
		fmt.Println("failt to update")
		return err
	}
	return nil
}
