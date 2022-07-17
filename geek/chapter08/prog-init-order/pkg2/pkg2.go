package pkg2

import (
	"fmt"
	_ "github.com/bigwhite/prog-init-order/pkg3"
)


var (
	_ = condtInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func condtInitCheck() string {
	if c1 != "" {
		// c1不是空值，所以说明已经初始化了 TODO：c1的空值为啥是是空字符串？
		fmt.Println("main: const c1 has been initialized")
	} 
	if c1 != "" {
		// c1不是空值，所以说明已经初始化了 TODO：c1的空值为啥是是空字符串？
		fmt.Println("main: const c1 has been initialized")
	} 
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized \n", name)
	return name
}

func init() {
	fmt.Println("main: pkg2 init func invoked")
}


func main() {
	fmt.Println("main: pkg2 main func invoked")
}