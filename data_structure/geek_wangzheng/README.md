

### 数据结构与算法之美-极客时间

[数据结构与算法之美](https://time.geekbang.org/column/intro/100017301?tab=catalog)

王争 前 Google 工程师



### 学习记录




### go 语法知识点

1. append实现切片合并
- append的第二个参数,不定长位置参数
```go
// func append(slice []Type, elems ...Type) []Type
package main

import "fmt"

func main() {
	left := []int{1, 2, 3}
	right := []int{5, 6}
	// all := append(append(left, right...), 4...)  // 错误
	all := append(append(left, right...), []int{4}...)
	fmt.Println(all)
}
```
- 使用变量
```go
package main

import "fmt"

func main() {
    left := []int{1, 2, 3}
    right := []int{5, 6}
    middle := []int{4}
    all := append(append(left, right...), middle...)
    fmt.Println(all)
}
```


