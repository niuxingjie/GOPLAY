package main

import (
	"fmt"
)

func main() {
	var null *Node

	liming := Person{NO: 1, Name: "liming", Age: 20}
	wangwu := Person{NO: 2, Name: "wangwu", Age: 20}

	second_node := Node{Value: liming, Next: null}
	first_node := Node{Value: wangwu, Next: &second_node}

	linked_list := LinkedList{FirstNode: &first_node}

	if result := linked_list.Search(0); result != nil {
		fmt.Println(result.Name)
	}
	if result := linked_list.Search(2); result != nil {
		fmt.Println(result.Name)
	}

	zhaoliu := Person{NO: 3, Name: "zhaoliu", Age: 20}
	linked_list.Lpush(zhaoliu)

	if result := linked_list.Search(3); result != nil {
		fmt.Println(result.Name)
	}

	linked_list.Delete(3)
	linked_list.Delete(1)
	fmt.Println(linked_list.GetLength())
	fmt.Println(linked_list.FirstNode.Value.Name)
}

// 存储的数据
type Person struct {
	NO   int
	Name string
	Age  int
}

// 节点
// 作为缓存的数据，不应该能够修改所以Value不存指针
// 作为Next记录下一个Node，链表上Node的Next需要修改，所以存指针
type Node struct {
	Value Person
	Next  *Node
}

// 单向链表
// FirstNode 存指针，才能修改Node的Next，达到更新链表本身数据目的。
type LinkedList struct {
	FirstNode *Node
}

// 查找
func (linked_list LinkedList) Search(no int) *Person {
	currentnode := linked_list.FirstNode
	if currentnode.Value.NO == no {
		return &currentnode.Value
	}
	for currentnode.Next != nil {
		currentnode = currentnode.Next
		if currentnode.Value.NO == no {
			return &currentnode.Value
		}
	}
	return nil
}

// 插入:头部插入O（1）,
// 由于不能计算出最后一个节点的内存地址只能轮询，尾部插入O（n）
func (linked_list *LinkedList) Lpush(person Person) {
	currentnode := linked_list.FirstNode
	first_code := Node{Value: person, Next: currentnode}
	linked_list.FirstNode = &first_code
}

// 删除
func (linked_list *LinkedList) Delete(no int) bool {
	currentnode := linked_list.FirstNode
	if currentnode.Value.NO == no {
		linked_list.FirstNode = currentnode.Next
		return true
	}
	for currentnode.Next != nil {
		if currentnode.Next.Value.NO == no {
			currentnode.Next = currentnode.Next.Next
		}
	}
	return false
}

// 删除
func (linked_list *LinkedList) GetLength() int {
	currentnode := linked_list.FirstNode
	length := 1
	if currentnode == nil {
		length := 0
		return length
	}
	for currentnode.Next != nil {
		currentnode = currentnode.Next
		length += 1
	}
	return length
}
