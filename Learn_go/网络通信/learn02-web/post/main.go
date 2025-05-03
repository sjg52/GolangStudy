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
	path := "http://apis.juhe.cn/simpleWeather/query"
	values := url.Values{}
	values.Set("city", "深圳")
	values.Set("key", "1234567890")
	/*	r, _ := http.PostForm(path, values)
		defer r.Body.Close()
		b, _ := io.ReadAll(r.Body)
		fmt.Printf("%v\n", string(b))*/

	r, err := http.PostForm(path, values)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Printf("%v\n", string(b))
}
