package models

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type Translate struct {
}

func GetTranslateKey(lang, setction, key string) string {
	fmt.Println(lang)

	langKeys, _ := config.NewConfig("ini", "conf/"+lang+".ini")
	langKeys.String("key")

	val := langKeys.String(setction + "::" + key)
	return val
}
