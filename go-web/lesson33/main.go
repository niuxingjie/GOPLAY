package main

import (
	"fmt"
)

type AnimalsAction interface {
	Jiao()
}

type Person struct {
	name string
	age  int
}

func (p Person) Jiao() {
	fmt.Println("name:", p.name, "jiao!")
}

func main1() {
	// var ming interface{}

	ming := Person{name: "ming", age: 50}

	ming.Jiao()
}

func main2() {
	var ming interface{}

	ming = Person{name: "ming", age: 50}

	ming.Jiao() //  (type interface{} has no field or method Jiao)
}

func main3() {
	var ming interface{}

	ming = Person{name: "ming", age: 50}

	// ming.Jiao() //  (type interface{} has no field or method Jiao)

	value, ok := ming.(Person) // 转换类型；也就是说可以属于某类型，和等于某类型是不一样的。
	if ok {
		value.Jiao()
	} else {
		fmt.Println(ming, "is not a type of Person")
	}
}

func main() {
	main1()
	main2()
	main3()
}
