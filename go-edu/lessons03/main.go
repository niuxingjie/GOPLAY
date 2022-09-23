package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("《Go语言极简一本通》")
	fmt.Println("公众号：面向加薪学习")
	pid := os.Getgid()
	fmt.Println(pid)
}
