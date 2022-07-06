package main

import "fmt"

func main() {
    // 变量名 := 类型 初始值
    countryCapitalMap := map[string]string {"France":"Paris", "Italy":"Rome","Japan":"Tokyo", "India":"New Delhi"}
    
    fmt.Println("打印原始map")
    printMap(countryCapitalMap)

    // 使用内置函数delete通过key删除
    // delete(countryCapitalMap, 'France')  // more than one character in rune literal  TODO: 单引号和双引号含义不一样？
    delete(countryCapitalMap, "France")

    fmt.Println("打印新map")
    printMap(countryCapitalMap)
}

func printMap(country_map map[string]string){
    for country := range country_map {
        fmt.Printf("key is %s, value is %s.\n", country, country_map[country])
    }
}