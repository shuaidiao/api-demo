package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id   int				`json:"id"`
	Name string 			`json:"name"`
	CreateTime time.Time	`json:"createTime" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time 	`json:"updateTime" orm:"auto_now;type(datetime)"`
}

type Response struct {
	Id   int			`json:"id"`
	Name string 		`json:"name"`
	CreateTime string	`json:"createTime"`
	UpdateTime string	`json:"updateTime"`
}

type Data struct {
	Records []Response `json:"records"`
	AllCount int64		`json:"allCount"`
}

type Select struct {
	Id int
}

type Add struct {
	Name string
}

type Update struct {
	Id int
	Name string
}

type Delete struct {
	Ids []int
}

type List struct {
	Name string
	Page int
	PageSize int
}

func UserSelect(userSelect *Select) (r *Response) {
	o := orm.NewOrm()
	var user User
	o.QueryTable("user").Filter("id", userSelect.Id).One(&user)
	var response Response
	response.Id = user.Id
	response.Name = user.Name
	response.CreateTime = user.CreateTime.Format("2006-01-02 15:04:05")
	response.UpdateTime = user.UpdateTime.Format("2006-01-02 15:04:05")
	return &response
}

func UserList(userList *List) (d *Data) {
	o := orm.NewOrm()
	var user []User
	qs := o.QueryTable("user").Filter("name__icontains", userList.Name)
	qs.Limit(userList.PageSize, (userList.Page-1)*userList.PageSize).All(&user)
	var responseList []Response
	for i := 0; i < len(user); i++ {
		var response Response
		response.Id = user[i].Id
		response.Name = user[i].Name
		response.CreateTime = user[i].CreateTime.Format("2006-01-02 15:04:05")
		response.UpdateTime = user[i].UpdateTime.Format("2006-01-02 15:04:05")
		responseList = append(responseList, response)
	}
	var data Data
	data.Records = responseList
	allCount,_ := qs.Count()
	data.AllCount = allCount
	return &data
}

func UserAdd(userAdd *Add) int {
	o := orm.NewOrm()
	var user User
	user.Name = userAdd.Name
	id,err := o.Insert(&user)
	if err == nil {
		return int(id)
	}
	return 0
}

func UserUpdate(userUpdate *Update) int {
	o := orm.NewOrm()
	var user User
	user.Id = userUpdate.Id
	o.Read(&user)
	user.Name = userUpdate.Name
	count,err := o.Update(&user)
	if err == nil {
		return int(count)
	}
	return 0
}

func UserDelete(userDelete *Delete) int {
	o := orm.NewOrm()
	count := 0
	for i := 0; i < len(userDelete.Ids); i++  {
		_,err := o.Delete(&User{Id: userDelete.Ids[i]})
		if err == nil {
			count++
		}
	}
	return count
}

