// 参考文档：https://blog.csdn.net/u012124304/article/details/109401561
package main

import "fmt"

type Animal interface {
	GetName() string
}

// Cat 实现 Animal 接口
type Cat struct {
	name string
}

func (c *Cat) GetName() string {
	return "I'm cat : " + c.name
}

// Dog 实现 Animal 接口
type Dog struct {
	name string
}

func (d *Dog) GetName() string {
	return "I'm dog : " + d.name
}

func main() {
	cat := Cat{
		name: "hello kitty",
	}

	animal := Animal(&cat) // 结构体转接口，括号里需要传递一个 *Cat 类型而不能是 Cat 类型，因为是 *Cat 类型实现了 GetName() 方法，而不是 Cat 类型
	fmt.Println(animal.GetName())

	dog1 := Dog(cat) // 结构体之间进行转换，括号里需要传递一个 Cat 类型，因为 Cat = Dog, *Cat = *Dog
	fmt.Println(dog1.GetName())

	dog2 := (*Dog)(&cat)
	fmt.Println(dog2.GetName()) // 如上所述，*Cat = *Dog

	cat2, ok := animal.(*Cat) // 类型断言，左边必须是一个接口类型的对象，当接口对象的实际类型和要转换的目标类型匹配时，转换成功，否则失败
	if ok {
		fmt.Println("convert animal to cat - " + cat2.GetName())
	} else {
		fmt.Println("can not convert animal to cat")
	}

	dog3, ok := animal.(*Dog) // 类型断言，接口对象的实际类型和要转换的目标类型不匹配
	if ok {
		fmt.Println("convert animal to dog - " + dog3.GetName())
	} else {
		fmt.Println("can not convert animal to dog")
	}

}
