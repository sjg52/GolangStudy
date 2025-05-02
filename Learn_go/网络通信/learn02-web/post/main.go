package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 发送post请求
func main() {
	postA()
}

func postA() {
	path := "http://apis.juhe.cn/simpleWeather/quary"
	values := url.Values{}
	values.Set("city", "深圳")
	r, _ := http.PostForm(path, values)
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	fmt.Printf("%v\n", string(b))
}
