package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	// 获取命令行参数
	var userID string
	var password string

	flag.StringVar(&userID, "u", "", "学号")
	flag.StringVar(&password, "p", "", "密码")
	flag.Parse()
	if userID == "" {
		fmt.Println("请使用-u 指明你的学号")
		os.Exit(1)
	}
	if password == "" {
		fmt.Println("请使用-p 指明你的密码")
		os.Exit(1)
	}
	//先随便发送一个请求，查看是否触发了重定向，判断是否是要求重新登陆的重定向 "获取在线情况" 百度的唯一用法，笑
	client := &http.Client{}
	resp, err := client.Get("http://baidu.com/")
	if err != nil {
		log.Panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err.Error())
	}
	sbody := string(body)
	// fmt.Println("header")
	// for key, value := range resp.Header {
	// 	fmt.Printf("KEY:%s  VALUE:%s \n", key, value)
	// }

	// fmt.Println("body", string(body))
	//判断是否需要认证，“如果包含 top.self.location.href就当作需要认证”
	if strings.Contains(sbody, "top.self.location.href") {
		//需要认证
		log.Println("检测到需要重新认证")
		//过滤出地址和设备id
		r, err := regexp.Compile("href='(.*)/eportal")
		handleErr(err)
		address := r.FindStringSubmatch(sbody)[1]
		r, err = regexp.Compile("index.jsp\\?(.*)'</script>")
		handleErr(err)
		device := r.FindStringSubmatch(sbody)[1]
		log.Println("地址", address)
		log.Println("设备", device)
		//然后构造post请求
		data := url.Values{}
		data.Set("userId", userID)
		data.Set("password", password)
		data.Set("queryString", device)
		body := strings.NewReader(data.Encode())
		client = &http.Client{}
		req, err := http.NewRequest("POST", address+"/eportal/InterFace.do?method=login", body)
		handleErr(err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
		resp, err := client.Do(req)
		handleErr(err)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		handleErr(err)
		m := make(map[string]interface{})
		err = json.Unmarshal(content, &m)
		handleErr(err)
		log.Println(m["result"], m["message"])
	} else {
		log.Println("不需要重新认证")
	}
}

func handleErr(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
