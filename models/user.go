package models

import (
	"database/sql"
	"fmt"

	//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	Id    int
	Name  string
	Age   int
	Likes string
}

func (u *UserModel) GetAllUser(dbName, table string) ([]UserModel, error) {
	db, err := GetConnDB(dbName)
	CheckError(err)

	sqlStr := fmt.Sprintf("select * from %s", table)
	rows, err := db.Query(sqlStr)
	CheckError(err)

	users := make([]UserModel, 1)
	//make([]interface{}, len(columns))

	for rows.Next() {
		var id, age int
		var name, likes string
		err := rows.Scan(&id, &name, &age, &likes)
		if err != nil {
			return nil, err
		}

		users_tmp := UserModel{Id: id, Age: age, Name: name, Likes: likes}
		users = append(users, users_tmp)
	}

	return users, nil
}

func GetConnDB(dbName string) (*sql.DB, error) {
	DbConfigInfo := GetDbConfig()
	db, err := sql.Open("mysql", DbConfigInfo["User"]+":"+DbConfigInfo["Pwd"]+"@/"+DbConfigInfo["DbName"]+"?charset=utf8")
	//db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	CheckError(err)
	return db, err
}

func (u *UserModel) AddUser() int64 {
	db, err := GetConnDB("test")
	CheckError(err)

	stmt, err := db.Prepare(`INSERT users(name,age,likes) values (?,?,?)`)
	defer stmt.Close()
	CheckError(err)

	res, err := stmt.Exec("cc", 20, "book,food,apple")
	CheckError(err)

	id, err := res.LastInsertId()
	CheckError(err)

	return id
}

//删除数据
func DeleteUser() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	CheckError(err)

	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	CheckError(err)
	res, err := stmt.Exec(1)
	CheckError(err)
	num, err := res.RowsAffected()
	CheckError(err)
	fmt.Println(num)
}

//更新数据
func update() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	CheckError(err)

	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	CheckError(err)
	res, err := stmt.Exec(21, 2, 1)
	CheckError(err)
	num, err := res.RowsAffected()
	CheckError(err)
	fmt.Println(num)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
		//fmt.Println(err)
	}
}

func GetDbConfig() map[string]string {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic("get app config is wrong")
	}

	DbConfigInfo := make(map[string]string)
	DbConfigInfo["DbHost"] = iniconf.String("mysqlurls")
	DbConfigInfo["User"] = iniconf.String("mysqluser")
	DbConfigInfo["Pwd"] = iniconf.String("mysqlpass")
	DbConfigInfo["DbName"] = iniconf.String("mysqldb")

	return DbConfigInfo
}
