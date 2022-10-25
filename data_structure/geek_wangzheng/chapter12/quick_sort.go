/*
快速排序：
- 1 先取出数列中的第一个数作为基准数
- 2 将剩余的数列中比基准数大的元素放在它的右边，比基准数小的放在它的左边。(此时，基础数的下标索引已经就找到了。)
- 3 然后再对左右两部分重复第二步操作，直到各区间只有一个数。

*/
package main

import (
	"fmt"
)

func main() {
	// 初始数组序列
	var a []int = []int{6, 9, 3, 4, 5, 1, 7, 2, 8, 0, 11, 15, 20}

	quick_sort(a, 0, len(a)-1)

	print_array(a)
}

// 打印数组
func print_array(a []int) {
	fmt.Println(a)
}

// 分区
func Partition(a []int, p int, r int) int {
	pivot := a[r]
	i := p
	for j := p; j <= r; j++ {
		if a[j] < pivot {
			a[j], a[i] = a[i], a[j]
			i = i + 1
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

// 排序数组
func quick_sort(a []int, p int, r int) {
	if p >= r {
		return
	}
	q := Partition(a, p, r)
	quick_sort(a, p, q-1)
	quick_sort(a, q+1, r)
}
