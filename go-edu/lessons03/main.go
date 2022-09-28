package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("《Go语言极简一本通》")
	fmt.Println("公众号：面向加薪学习")
	pid := os.Getgid() // https://pkg.go.dev/os@go1.19.1#Getgid
	fmt.Println(pid)
}
