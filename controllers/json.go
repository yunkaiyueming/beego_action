package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type JsonTestController struct {
	beego.Controller
}

type msg struct {
	Id    int
	Name  string
	Likes []string
}

//各种json数据测试
func (j *JsonTestController) Get() {
	TestSlice := []int{1, 2, 4, 5, 6}
	test_json, _ := json.Marshal(TestSlice)
	fmt.Println(string(test_json))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	msg_info := msg{Id: 1, Name: "aa", Likes: []string{"appale", "banana", "orange"}}
	msg_infos := make([]msg, 10, 10)
	for i := 0; i < 10; i++ {
		msg_infos = append(msg_infos, msg_info)
	}
	msg_info_json, _ := json.Marshal(msg_infos)
	fmt.Println(string(msg_info_json))
}
