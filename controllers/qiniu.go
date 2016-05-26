package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/kodo"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodocli"
)

type QiniuController struct {
	beego.Controller
}

var (
	//设置需要操作的空间
	bucket = "game-package-world-war"
	//设置需要操作的文件的key
	key        = "icon512.png"
	domain     = "xxxx.com2.z0.glb.qiniucdn.com"
	access_key = ""
	select_key = ""
)

func (q *QiniuController) GetFilesMsg() {
	conf.ACCESS_KEY = access_key
	conf.SECRET_KEY = select_key

	//new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	//调用Stat方法获取文件的信息
	entry, err := p.Stat(nil, key)
	//打印出错时返回的信息
	if err != nil {
		fmt.Println(err)
	}

	file_msg, _ := json.Marshal(entry)
	fmt.Println(string(file_msg))
	q.Data["json"] = entry
	q.ServeJSON()
}

func (q *QiniuController) GetDownFileUrl() {
	down_fiel_url := downloadUrl(domain, key)
	fmt.Println(down_fiel_url)
}

//调用封装好的downloadUrl方法生成一个下载链接
func downloadUrl(domain, key string) string {
	//调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := kodo.MakeBaseUrl(domain, key)
	policy := kodo.GetPolicy{}
	//生成一个client对象
	c := kodo.New(0, nil)
	//调用MakePrivateUrl方法返回url
	return c.MakePrivateUrl(baseUrl, &policy)
}

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func (q *QiniuController) SimpleUploadFile() {
	filepath := "E:/www2/GitHub/gitbook_test/README.md"
	simpleUploadFile(filepath)
}

func simpleUploadFile(filepath string) {
	conf.ACCESS_KEY = access_key
	conf.SECRET_KEY = select_key

	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		SaveKey: filepath,
		Scope:   bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)
	fmt.Println(token)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置上传文件的路径

	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}
