package controllers

import (
	"beego_action/helpers"
	"beego_action/models"
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type UserController struct {
	BaseController
}

func (this *UserController) Construct() {
	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"
}

func (this *UserController) MyRender(viewFile string) {
	this.Construct()
	this.BaseController.MyRender(viewFile)
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

func (this *UserController) Nginx() {
	redis_con := helpers.GetCon()

	log_keys, err := redis.Strings(redis_con.Do("keys", "access_log:*"))
	if err != nil {
		fmt.Println(err)
	}

	log_infos := make(map[string]string, 200)
	for _, log_key := range log_keys {
		log_item_str, _ := redis.Strings(redis_con.Do("hGetAll", log_key))
		log_infos[log_key] = strings.Join(log_item_str, "; ")

		if len(log_infos) > 30 {
			break
		}
	}

	item_descs := []string{
		"client_ip",
		"request_time",
		"request_url",
		"status",
		"body_bytes_sent",
		"http_referer",
		"user_agent",
		"http_x_forwarded_for",
	}

	this.Data["log_infos"] = log_infos
	this.Data["item_descs"] = item_descs

	this.MyRender("user/view_nginx.html")
}
