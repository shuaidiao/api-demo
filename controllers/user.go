package controllers

import (
	"api-demo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Select() {
	var userSelect models.Select
	json.Unmarshal(u.Ctx.Input.RequestBody, &userSelect)
	fmt.Println(userSelect)
	valid := validation.Validation{}
	valid.Required(userSelect.Id, "id")
	if valid.HasErrors() {
		u.Data["json"] = Response{100, "参数错误", nil}
		u.ServeJSON()
		return
	}
	u.Data["json"] = Response{200, "成功", models.UserSelect(&userSelect)}
	u.ServeJSON()
}

func (u *UserController) List() {
	var userList models.List
	json.Unmarshal(u.Ctx.Input.RequestBody, &userList)
	fmt.Println(userList)
	valid := validation.Validation{}
	valid.Required(userList.Page, "page")
	valid.Required(userList.Page, "pageSize")
	if valid.HasErrors() {
		u.Data["json"] = Response{100, "参数错误", nil}
		u.ServeJSON()
		return
	}
	u.Data["json"] = Response{200, "成功", models.UserList(&userList)}
	u.ServeJSON()
}

func (u *UserController) Add() {
	var userAdd models.Add
	json.Unmarshal(u.Ctx.Input.RequestBody, &userAdd)
	fmt.Println(userAdd)
	valid := validation.Validation{}
	valid.Required(userAdd.Name, "name")
	if valid.HasErrors() {
		u.Data["json"] = Response{100, "参数错误", nil}
		u.ServeJSON()
		return
	}
	count := models.UserAdd(&userAdd)
	if count < 1 {
		u.Data["json"] = Response{100, "添加失败", nil}
		u.ServeJSON()
		return
	}
	u.Data["json"] = Response{200, "成功", nil}
	u.ServeJSON()
}

func (u *UserController) Update() {
	var userUpdate models.Update
	json.Unmarshal(u.Ctx.Input.RequestBody, &userUpdate)
	fmt.Println(userUpdate)
	valid := validation.Validation{}
	valid.Required(userUpdate.Name, "name")
	if valid.HasErrors() {
		u.Data["json"] = Response{100, "参数错误", nil}
		u.ServeJSON()
		return
	}
	count := models.UserUpdate(&userUpdate)
	if count != 1 {
		u.Data["json"] = Response{100, "更新失败", nil}
		u.ServeJSON()
		return
	}
	u.Data["json"] = Response{200, "成功", nil}
	u.ServeJSON()
}

func (u *UserController) Delete() {
	var userDelete models.Delete
	json.Unmarshal(u.Ctx.Input.RequestBody, &userDelete)
	fmt.Println(userDelete)
	valid := validation.Validation{}
	valid.Required(userDelete.Ids, "ids")
	if valid.HasErrors() {
		u.Data["json"] = Response{100, "参数错误", nil}
		u.ServeJSON()
		return
	}
	count := models.UserDelete(&userDelete)
	if count < 1 {
		u.Data["json"] = Response{100, "删除失败", nil}
		u.ServeJSON()
		return
	}
	u.Data["json"] = Response{200, "成功", nil}
	u.ServeJSON()
}
