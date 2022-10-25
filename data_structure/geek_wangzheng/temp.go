package main

import "fmt"

func main() {
	left := []int{1, 2, 3}
	right := []int{5, 6}
	// all := append(append(left, right...), 4...)  // 错误
	all := append(append(left, right...), []int{4}...)
	fmt.Println(all)
}
