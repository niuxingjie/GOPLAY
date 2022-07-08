package main

import (
    "fmt"
    "reflect"
)

// TODO: interface{}这里为什么有大括号？
func reflectsetvalue1(x interface{}) {
    value := reflect.ValueOf(x)
    if value.Kind() == reflect.String {
        value.SetString("欢迎来到W3Cschool")
    }
}

func reflectsetvalue2(x interface{}) {
    value := reflect.ValueOf(x)

    // 反射中使用Elem()方法获取指针所指向的值
    if value.Elem().Kind() == reflect.String {
        value.Elem().SetString("欢迎来到W3Cschool")
    }
}


func main(){
    address := "www.w3cshool.cn"
    // reflectsetvalue1(address)  // panic: reflect: reflect.Value.SetString using unaddressable value
    fmt.Println(address)

    reflectsetvalue1(&address)
    fmt.Println(address)  // www.w3cshool.cn

    reflectsetvalue2(&address)
    fmt.Println(address)  // 欢迎来到W3Cschool
}