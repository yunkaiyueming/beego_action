package models

import (
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

func init() {
	fmt.Println("start user_model init")
}

func (u *UserModel) GetAllUser() ([]UserModel, error) {
	sqlStr := fmt.Sprintf("select * from %s", USER_MODEL_TABLE_NAME)
	fmt.Println(sqlStr)
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
	stmt, err := db.Prepare(`INSERT users(name,age,likes,pwd) values (?,?,?,?)`)
	defer stmt.Close()
	CheckError(err)

	res, err := stmt.Exec("cc", 20, "book,food,apple", "")
	CheckError(err)

	id, err := res.LastInsertId()
	CheckError(err)

	return id
}

//删除数据
func (u *UserModel) DeleteUserById(id int) {
	stmt, err := db.Prepare(`DELETE FROM users WHERE id=?`)
	CheckError(err)
	res, err := stmt.Exec(id)
	CheckError(err)
	num, err := res.RowsAffected()
	CheckError(err)
	fmt.Println(num)
}

//更新数据
func (u *UserModel) UpdateUserById(id int, data map[string]string) {
	stmt, err := db.Prepare(`UPDATE users SET name=?,age=?,likes=? WHERE id=?`)
	CheckError(err)

	fmt.Println(data)
	res, err := stmt.Exec(data["name"], data["age"], data["likes"], id)
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

func (u *UserModel) GetUserById(id int) UserModel {
	sql_str := fmt.Sprintf("select * from %s where id='%d'", USER_MODEL_TABLE_NAME, id)
	fmt.Println(sql_str)
	row := db.QueryRow(sql_str)

	var age int
	var name, likes, pwd string
	row.Scan(&id, &name, &age, &likes, &pwd)

	return UserModel{Id: id, Age: age, Name: name, Likes: likes}
}
