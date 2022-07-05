package main

import "fmt"

func main(){
    var numbers = [5]int{1, 2, 3, 4, 5}
    var sum int

    for _, num := range numbers {
        sum += num
    }
    fmt.Printf("len=%d cap=%d sum=%v\n", len(numbers), cap(numbers), sum)

    for i, num := range numbers {
        fmt.Printf("第%d个元素：%d\n", i, num)
    }
}