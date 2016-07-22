package controllers

import (
	"beego_action/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	id, _ := this.GetInt("id")
	u := &models.UserModel{}
	userInfo := u.GetUserById(id)

	this.Data["userInfo"] = userInfo
	this.MyRender("user/view_index.html")
}

func (this *UserController) GetAllUser() {
	u := &models.UserModel{}
	users, _ := u.GetAllUser()

	this.Data["users"] = users
	this.MyRender("user/get_user.html")
}

func (this *UserController) UpdateUser() {
	id, _ := this.GetInt("id")
	action := this.GetString("action")
	u := &models.UserModel{}
	if action == "" {
		userInfo := u.GetUserById(id)
		fmt.Println(userInfo)

		this.Data["userInfo"] = userInfo
		this.MyRender("user/view_index.html")
	} else {
		u := &models.UserModel{}
		name := this.GetString("name")
		age := this.GetString("age")
		likes := this.GetString("likes")

		data := map[string]string{"name": name, "age": age, "likes": likes}
		u.UpdateUserById(id, data)

		this.Ctx.Redirect(200, "http://localhost:8080/user/index")
	}
}

func (this *UserController) AddUser() {
	u := &models.UserModel{}
	id := u.AddUser()
	if id > 0 {
		msg := fmt.Sprintf("add user success %d", id)
		this.Ctx.WriteString(msg)
	} else {
		this.Ctx.WriteString("add user false")
	}
}

func (this *UserController) DeleteUser() {
	id, _ := this.GetInt("id")
	fmt.Println(id)

	u := &models.UserModel{}
	u.DeleteUserById(id)
	this.Ctx.WriteString("hello deleteuser")
}
