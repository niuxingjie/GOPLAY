
### 从零开始学 Go Web 编程

[从零开始学 Go Web 编程：build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)

[go 在线工具](https://c.runoob.com/compile/21/)

### 要点记录

1. 类型转换及断言

- 接口变量及持有的变量。
```
接口(interface)变量可用于存储符合接口(interface)的任何值，并调用属于该接口(interface)的方法。
请注意，您将无法通过接口(interface)变量访问基础值上的字段。
您可以通过类型断言提取接口(interface)变量持有的动态值。
```

- code
```go
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
```


1. 防止多次递交表单
```
https://github.com/niuxingjie/build-web-application-with-golang/blob/master/zh/04.4.md

解决方案是在表单中添加一个带有唯一值的隐藏字段。
在验证表单时，先检查带有该唯一值的表单是否已经递交过了。
如果是，拒绝再次递交；如果不是，则处理表单进行逻辑处理。
另外，如果是采用了Ajax模式递交表单的话，当表单递交后，通过javascript来禁用表单的递交按钮。
```

2. 位运算符
```go
package main

import "fmt"

func main() {
	fmt.Println("2的10次方就是1024=4*4*4*4*4=", 4*4*4*4*4)
    fmt.Println("n<<10就是n*1024：", 32<<20/1024)
    fmt.Println("n<<20就是n*1024*1024：", 32<<20/1024/1024)
	fmt.Println("32<<10 B就可以表示32KB")
	fmt.Println("32<<20 B就可以表示32MB")
}

// 2的10次方就是1024=4*4*4*4*4= 1024
// n<<10就是n*1024： 32768
// n<<20就是n*1024*1024： 32
// 32<<10 B就可以表示32KB
// 32<<20 B就可以表示32MB
```
