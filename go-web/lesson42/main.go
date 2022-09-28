/*
4.2 验证表单的输入
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world!")
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	fmt.Println("method:", method) //获取请求的方法
	switch method {
	case "GET":
		fmt.Fprintf(w, "请使用POST传递username和password")
		return
	case "POST":
		// 注意Content-Type使用application/x-www-form-urlencoded
		err := r.ParseForm()
		//请求的是登录数据，那么执行登录的逻辑判断
		if err != nil {
			fmt.Fprintf(w, "请使用POST传递username和password")
		}
		fmt.Println("username:", r.PostForm["username"])
		fmt.Println("password:", r.PostForm["password"])
	default:
		fmt.Fprintf(w, "请使用POST传递username和password")
		return
	}

}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
