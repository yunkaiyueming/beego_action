package controllers

import (
	"beego_action/helpers"
	"fmt"
	"time"
)

type TimeController struct {
	BaseController
}

func (this *TimeController) Test() {
	//Cron()
	//Cron2()
	//Cron3()
	TestMail()
}

func Cron() {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range ticker.C {
			Task("1:")
		}
	}()
}

func Cron2() {
	for {
		time.Sleep(time.Second * 2) //time.Sleep函数，表示休眠多少时间，休眠时处于阻塞状态，后续程序无法执行．
		Task("2:")
	}
}

func Cron3() {
	for {
		Task("3:")
		//time.After函数， 表示多少时间之后，但是在取出channel内容之前不阻塞，后续程序可以继续执行
		<-time.After(time.Second * 3)
	}
}

func Task(str string) {
	fmt.Println(str, time.Now())
}

//实现了定时任务
func demo(input chan interface{}) {
	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)
	msg := make(chan interface{})

	for {
		select {
		case msg <- input:
			println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			println("10s timer")
			t2.Reset(time.Second * 10) //重置一个10s的定时器
		}
	}
}

//没有实现定时任务
func demo2(input chan interface{}) {
	msg := make(chan interface{})
	for {
		select {
		case msg <- input:
			println(msg)

		case <-time.After(time.Second * 5):
			println("5s timer")

		case <-time.After(time.Second * 10): //永远不会执行，5s的到了，都是启动了一个新的从头开始计时的 Timer 对象
			println("10s timer")
		}
	}
}

func TestMail() {
	user := "from_user"
	password := "****"
	host := "smtp.exmail.qq.com:25"
	to := "to_user"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	err := helpers.SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println(err)
	}
}
