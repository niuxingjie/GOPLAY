package main

import (
	"fmt"
)

func main() {
	fmt.Println("结构体的类型判断：")

	liming := Person{Name: "liming", Age: 28}

	if p, ok := liming.(Person); ok {
		p.Run()
		p.Eat()
	}

	switch liming.(type) {
	case PersonAction:
		liming.Run()
		liming.Eat()
	default:
		fmt.Println(liming.Name, "不属于PersonAction类型！")

	}
}

type PersonAction interface {
	Run()
	Eat()
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Run() {
	fmt.Println(p.Name, " can run")
}

func (p *Person) Eat() {
	fmt.Println(p.Name, " can eat!")
}
