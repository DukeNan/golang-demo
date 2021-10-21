package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getUrl() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))

}

func getUrlByParms() {
	apiUrl := "http://192.168.100.61:8010"
	data := url.Values{}
	data.Set("name", "laowang")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed ,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func postDemo() {
	url := "http://192.168.100.61:8010/ai_coach/homepage"
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "user_id=12345"
	// json
	// contentType := "application/json"
	// data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("getr resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}

func main() {
	// getUrl()
	// getUrlByParms()
	postDemo()
}
