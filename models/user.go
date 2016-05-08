package models

import (
	"database/sql"
	"fmt"
)

type UserModel struct {
	Id    int
	Name  string
	Age   int
	Likes string
	Pwd   string
}

const USER_MODEL_TABLE_NAME string = "users"

func (u *UserModel) GetAllUser() ([]UserModel, error) {
	sqlStr := fmt.Sprintf("select * from %s", USER_MODEL_TABLE_NAME)
	rows, err := db.Query(sqlStr)
	CheckError(err)

	users := make([]UserModel, 0)
	//make([]interface{}, len(columns))

	for rows.Next() {
		var id, age int
		var name, likes, pwd string
		err := rows.Scan(&id, &name, &age, &likes, &pwd)
		if err != nil {
			return nil, err
		}

		users_tmp := UserModel{Id: id, Age: age, Name: name, Likes: likes, Pwd: pwd}
		users = append(users, users_tmp)
	}

	return users, nil
}

func (u *UserModel) AddUser() int64 {
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
func UpdateUser() {
	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	CheckError(err)
	res, err := stmt.Exec(21, 2, 1)
	CheckError(err)
	num, err := res.RowsAffected()
	CheckError(err)
	fmt.Println(num)
}

func (u *UserModel) CheckGetUser(check_name, check_pwd string) UserModel {
	sql_str := fmt.Sprintf("select * from %s where name='%s' and pwd='%s'", USER_MODEL_TABLE_NAME, check_name, check_pwd)
	fmt.Println(sql_str)
	row := db.QueryRow(sql_str)

	var id, age int
	var name, likes, pwd string

	row.Scan(&id, &name, &age, &likes, &pwd)

	return UserModel{Id: id, Age: age, Name: name, Likes: likes, Pwd: pwd}
}
