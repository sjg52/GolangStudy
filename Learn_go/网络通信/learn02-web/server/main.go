package main

import (
	"fmt"
	"net/http"
)

// 服务器端
func main() {
	fmt.Printf("准备开启服务器…………")
	http.HandleFunc("/hello", hello)
	handler1 := handler1{}

	http.Handle("/helloA", &handler1)
	http.ListenAndServe(":8081", nil)
}

type handler1 struct{}

func (h1 *handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "handler1")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello,web")
}
