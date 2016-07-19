package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id           int
	Article_name string
	Contents     string
	Create_time  string
	Author       string
	Type         string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Article))

	// create table
	//orm.RunSyncdb("default", false, true)
}

func aa() {
	fmt.Println("aa")
}

func BB() {
	fmt.Println("bb")
}

var article Article = Article{Article_name: "slene", Contents: "hehe", Create_time: "2016-01-01"}

func (a *Article) Read() {
	fmt.Println("article control read")
	o := orm.NewOrm()
	// read one
	a2 := Article{Id: 1}
	err := o.Read(&a2)
	fmt.Printf("ERR: %v\n", err)
	fmt.Println(a2)
}

/*
func (a *Article) Insert() {
	o := orm.NewOrm()

	// insert
	id, err := o.Insert(&article)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

func (a *Article) Delete() {
	o := orm.NewOrm()
	// delete
	num, err := o.Delete(&article)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}



func (a *Article) Update() {
	o := orm.NewOrm()
	// update
	article.Article_name = "astaxie"

	num, err := o.Update(&article)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
*/
