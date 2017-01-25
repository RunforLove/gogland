package main

import (
	"net/http"
	"io/ioutil"
	"os"
)

// 定义结构体
type ThreadItem struct {
	url string
	content string
	imgs []string
}

// 定义函数，下载网页
func httpGet(url string) (content string, statuscode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statuscode = -100
		return
	}
	defer resp.Body.Close()
	data,err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statuscode = -200
		return
	}
	statuscode = resp.StatusCode
	content = string(data)
	return
}

