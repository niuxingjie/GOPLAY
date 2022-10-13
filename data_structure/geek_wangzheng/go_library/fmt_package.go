/*
[Go Print Statement](https://www.programiz.com/golang/print-statement)
*/
package main

import "fmt"

func main() {
	// 1 fmt.Print

	// 前后无换行打印
	fmt.Print("Hello ")
	fmt.Print("World!")

	// 打印变量
	name := "Li Ming"
	fmt.Print(name)

	// 打印多值,多值间无空格
	fmt.Print("Hello ", name)

	// 2 fmt.Println
	fmt.Println("line1")
	fmt.Println("line2")                  // 末尾有换行
	fmt.Println("line3", " Hello:", name) // 多值间有空格

	// 3 fmt.Printf
	fmt.Printf("a int = %d", 5) // 末尾无换行
	println()
	fmt.Printf("a float = %g", 5.5)
	println()
	fmt.Printf("a string = %s", "5.5") // 类型需要一致
	println()
	fmt.Printf("a bool = %t", true)
	println()

	age := 23
	fmt.Printf("My name is %s and i am %d years old!", name, age)

	print("...") // 等同于fmt.Print
	println()    // 等同于fmt.Println

	// 4 打印复杂类型数据
	array := [2]int{1, 2}
	fmt.Print(array)
	println()

	type Person struct {
		name string
		age  int
	}
	p := Person{
		name: name,
		age:  23,
	}
	fmt.Print(p)
	println()
}
