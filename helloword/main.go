package main

import "fmt"

func main(){
    var numbers = []int{1, 2, 3, 4, 5, 6}
    var i int
    var numbers_copy = numbers
    // 改变 numbers[0]、numbers[1]
    change_0(numbers)

    // 改变 numbers[2]、numbers[3]
    numbers_copy[2] = 0
    numbers_copy[3] = 0

    for i=0; i<6; i++{
        fmt.Printf("numbers[%d]的平均值为:%d\n", i, numbers[i])
    }
    
}

func change_0(array []int) {
    array[0] = 0
    array[1] = 0
}