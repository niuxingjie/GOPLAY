package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// 冒泡排序
	// 初始数组序列
	var a []int = []int{6, 9, 3, 4, 5, 1, 7, 2, 8, 0}

	test_sort(bubble_sort_1, a)

	a = []int{5, 1, 2, 3, 4}
	test_sort(bubble_sort_2, a)
}

func test_sort(f func([]int, int) []int, a []int) {
	fmt.Printf("测试%s\n", GetFunctionName(f))
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
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// 排序数组
func bubble_sort_1(a []int, n int) []int {
	// go 数组能直接赋值给切片参数吗:数组是值类型，切片是引用类型。
	fmt.Println("bubble_sort_1")
	for i := 0; i < n-1; i++ {
		fmt.Printf("第%d次冒泡\n", i+1)
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// 排序数组
func bubble_sort_2(a []int, n int) []int {
	// go 数组能直接赋值给切片参数吗:数组是值类型，切片是引用类型。
	fmt.Println("bubble_sort_2")
	flag := true
	for i := 0; i < n-1; i++ {
		if flag {
			fmt.Printf("第%d次冒泡\n", i+1)
			for j := 0; j < n-1; j++ {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
					flag = false
				}
			}
		} else {
			break
		}
	}
	return a
}
