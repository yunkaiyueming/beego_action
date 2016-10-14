package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
)

func My_http_get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		ret := map[string]string{"code": "500", "msg": "get error occur"}
		ret_json, _ := json.Marshal(ret)
		return string(ret_json)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ret := map[string]string{"code": "500", "msg": "body error occur"}
		ret_json, _ := json.Marshal(ret)
		return string(ret_json)
	}

	return string(body)
}

//post 第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递
func My_http_post() {
	resp, err := http.Post("http://192.168.7.180:9090/user_list", "application/x-www-form-urlencoded", strings.NewReader("username=aaa"))
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body error occur")
	}
	fmt.Println(string(body))
}

func HttpPostForm() {
	resp, err := http.PostForm("http://192.168.7.180:9090/user_list", url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

//有时需要在请求的时候设置头参数、cookie之类的数据,上面的post请求，必须要设定Content-Type为application/x-www-form-urlencoded，post参数才可正常传递
func HttpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://192.168.7.180:9090/user_list", strings.NewReader("username=aab"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func ClawResponseHeader(url string) http.Header {
	resp, err := http.Get(url)
	if err != nil {
		return nil
		//fmt.Println("error occur")
	}

	defer resp.Body.Close()
	response_headers := resp.Header
	return response_headers
	//	for key, val := range response_headers {
	//		fmt.Println(key, val)
	//	}
}

func GetAllResponse(url string) http.Response {
	resp, _ := http.Get(url)
	return *resp
}

func HttpGetDo(url string, req *http.Request) *http.Response {
	client := &http.Client{}
	real_req, _ := http.NewRequest("GET", url, nil)

	for key, val := range req.Header {
		real_req.Header.Set(key, val[0])
	}

	resp, _ := client.Do(real_req)
	return resp
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func SiteUrl(baseName string) string {
	return beego.AppConfig.String("siteUrl") + "/" + baseName
}
