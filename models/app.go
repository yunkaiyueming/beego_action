package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AppModel struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	dbConfig := GetDbConfig()
	dbhost := dbConfig["DbHost"]
	dbport := dbConfig["Port"]
	dbuser := dbConfig["User"]
	dbpassword := dbConfig["Pwd"]
	dbname := dbConfig["DbName"]
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn, 30)

	// register model
	orm.RegisterModel(new(AppModel), new(Article))
}

func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		fmt.Println(err)
	}
}

func (this *AppModel) TestCurd() {
	o := orm.NewOrm()

	user := AppModel{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := AppModel{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

func (this *AppModel) InsertMulitTest() {
	apps := []AppModel{
		{Id: 1, Name: "app1"},
		{Id: 2, Name: "app1"},
		{Id: 3, Name: "app1"},
		{Id: 4, Name: "app1"},
		{Id: 5, Name: "app1"},
	}

	o := orm.NewOrm() //注册新的orm
	o.InsertMulti(5, apps)
	fmt.Println("success")
}
