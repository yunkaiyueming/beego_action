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
	bucket     string
	domain     string
	access_key string
	select_key string
	key        = "README.md"
)

func init() {
	bucket = beego.AppConfig.String("bucket")
	domain = beego.AppConfig.String("domain")
	access_key = beego.AppConfig.String("qiniu_ak")
	select_key = beego.AppConfig.String("qiniu_sk")
}

//获取文件信息
func (this *QiniuController) GetFilesMsg() {
	conf.ACCESS_KEY = access_key
	conf.SECRET_KEY = select_key

	//new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	//调用Stat方法获取文件的信息
	key = "640.jpg"
	entry, err := p.Stat(nil, key)
	//打印出错时返回的信息
	if err != nil {
		fmt.Println(err)
	}

	file_msg, _ := json.Marshal(entry)
	fmt.Println(string(file_msg))
	this.Data["json"] = entry
	this.ServeJSON()
}

//获取下载URL
func (this *QiniuController) GetDownFileUrl() {
	down_fiel_url := downloadUrl(domain, key)
	fmt.Println(down_fiel_url)
	this.Data["json"] = down_fiel_url
	this.ServeJSON()
}

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

//简单上传文件
func (this *QiniuController) SimpleUploadFile() {
	filepath := "E:/GO_PATH/src/bit-every-day/README.md"
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

func (this *QiniuController) GetBucketFiles() {
	conf.ACCESS_KEY = access_key
	conf.SECRET_KEY = select_key

	//new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	//调用List方法，第二个参数是前缀,第三个参数是delimiter,第四个参数是marker，第五个参数是列举条数
	ListItem, _, _, _ := p.List(nil, "", "", "", 100)

	//循环遍历每个操作的返回结果
	for _, item := range ListItem {
		fmt.Println(item.Key, item.Fsize)
	}
	this.Data["json"] = ListItem
	this.ServeJSON()
}
