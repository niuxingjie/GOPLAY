package main

import (
	"fmt"
)

func main() {
	fmt.Println("结构体的类型判断：")
	liming := Person{Name: "liming", Age: 28}
	Skill(liming)
}

type PersonAction interface {
	Run()
	Eat()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Run() {
	fmt.Println(p.Name, " can run")
}

func (p Person) Eat() {
	fmt.Println(p.Name, " can eat!")
}

// 接口的静态特性：PersonAction持有实现了接口方法的结构体
// 编写Skil(Person{Name: "liming", Age: 20})是可以通过编译的
func Skill(p PersonAction) {
	p.Run()
	p.Eat()
}
