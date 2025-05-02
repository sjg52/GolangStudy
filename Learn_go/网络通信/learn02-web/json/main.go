package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	testGetJson()
}

func testGetJson() {
	path := "https://e.weather.com.cn/data/cityinfo/101010100.html"
	r, _ := http.Get(path)
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	fmt.Printf("%v\n", string(b))
	var jsonStr res
	json.Unmarshal([]byte(b), &jsonStr)
	fmt.Printf("%v\n", jsonStr)
}

type res struct {
	Info weather `json:"weatherinfo"`
}
type weather struct {
	City    string `json:"city"`
	Cityid  string `json:"cityid"`
	Temp1   string `json:"temp1"`
	Temp2   string `json:"temp2"`
	Weather string `json:"weather"`
	Img1    string `json:"img1"`
	Img2    string `json:"img2"`
	Ptime   string `json:"ptime"`
}
