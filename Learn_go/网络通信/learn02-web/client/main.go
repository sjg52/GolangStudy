package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	/*resp, _ := http.Get("https://www.baidu.com/")
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%v\n", string(b))*/

	params := url.Values{}
	params.Set("t", "govall")
	params.Set("q", "百度热搜")
	Url, err := url.Parse("https://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
	}
	Url.RawQuery = params.Encode()

	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println(err)
	}
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%v\n", string(b))

}
