package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var (
	redis_http string
	redis_port string
	redis_auth string
)

func init() {
	redis_http = beego.AppConfig.String("redis_http")
	redis_port = beego.AppConfig.String("redis_port")
	redis_auth = beego.AppConfig.String("redis_auth")
}

func GetCon() redis.Conn {
	c, err := redis.Dial("tcp", redis_http+":"+redis_port, redis.DialPassword(redis_auth))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return c
}

func DestructCon() {
	//defer redis.Close()
}
