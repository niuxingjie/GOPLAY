/*
插入排序

*/
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// 初始数组序列
	var a []int = []int{6, 9, 3, 4, 5, 1, 7, 2, 8, 0}

	test_sort(insert_sort, a)

	a = []int{2, 1, 3, 4, 5, 6}

	test_sort(insert_sort, a)
}

func test_sort(f func([]int, int) []int, a []int) {
	fmt.Printf("测试%s:", GetFunctionName(f))
	// 打印初始数据
	print_array(a)
	// 排序数据
	a = f(a, len(a))
	// 打印新数组
	print_array(a)
}

// 打印数组
func print_array(a []int) {
	fmt.Println(a)
}

// 打印函数名
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// 排序数组
func insert_sort(a []int, n int) []int {
	// 处理边界情况
	if n <= 1 {
		return a
	}
	times := 0
	fmt.Println("insert_sort")
	for i := 0; i < n-1; i++ {
		j := i + 1
		if a[i] > a[j] {
			times++
			fmt.Printf("第%d次插入\n", times)
			for ; j > 0; j-- {
				if a[j] < a[j-1] {
					a[j], a[j-1] = a[j-1], a[j]
				} else {
					break
				}
			}
		}
	}
	return a
}
