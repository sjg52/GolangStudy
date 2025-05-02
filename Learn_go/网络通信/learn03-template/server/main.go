package main

import (
	"html/template"
	"net/http"
)

func main() {
	/*name := "蜘蛛子"
	myTemplate := "hello,{{.}}"
	tmpl, err := template.New("test").Parse(myTemplate)
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		fmt.Println(err)
	}*/

	/*preson := Person{"蜘蛛子", 16}
	myTemplateA := "hello,{{.Name}},your age is {{.Age}}\n"
	tmpl, err := template.New("test").Parse(myTemplateA)
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(os.Stdout, preson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("准备开启服务器………")*/
	handler1 := handler1{}
	http.Handle("/handler1", &handler1)

	s := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: nil,
	}
	s.ListenAndServe()
}

type handler1 struct{}

func (h *handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("F:/Go/Learn_go/网络通信/learn03-template/html/index.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "Hello MyHTML")
}

type Person struct {
	Name string
	Age  int
}
