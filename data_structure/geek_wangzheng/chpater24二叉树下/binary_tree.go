package main

import (
	"fmt"
)

func main() {
	l1 := Person{NO: 1, Name: "L1", Age: 20}
	l2 := Person{NO: 2, Name: "L2", Age: 20}
	l3 := Person{NO: 3, Name: "L3", Age: 20}
	l4 := Person{NO: 4, Name: "L4", Age: 20}
	l5 := Person{NO: 5, Name: "L5", Age: 20}
	l6 := Person{NO: 6, Name: "L6", Age: 20}

	var binary_tree *BinaryTree = nil
	binary_tree = binary_tree.Add(Node{Value: l5})
	binary_tree.Add(Node{Value: l3})
	binary_tree.Add(Node{Value: l6})
	binary_tree.Add(Node{Value: l2})
	binary_tree.Add(Node{Value: l4})
	binary_tree.Add(Node{Value: l1})
	Print(binary_tree.TopNode)
}

type Person struct {
	NO   int
	Name string
	Age  int
}

type Node struct {
	Value Person
	Left  *Node
	Right *Node
}

// 二叉查找树
type BinaryTree struct {
	TopNode *Node
}

// 二叉查找树新增节点
func (tree *BinaryTree) Add(node Node) *BinaryTree {
	if tree == nil || tree.TopNode == nil {
		tree = &BinaryTree{TopNode: &node}
	} else {
		tree.TopNode.Add(node)
	}
	return tree
}

// 新增节点：找到合适的位置插入
func (current_node *Node) Add(node Node) bool {
	if current_node.Value.NO == node.Value.NO {
		return false
	} else if current_node.Value.NO > node.Value.NO {
		if current_node.Left == nil {
			current_node.Left = &node
			return true
		} else {
			return current_node.Left.Add(node)
		}
	} else {
		if current_node.Right == nil {
			current_node.Right = &node
			return true
		} else {
			return current_node.Right.Add(node)
		}
	}
}

// 前序遍历查找：当前节点 --> 左节点 --> 右节点
// 从二叉树的顶点查找节点：找到节点返回
func (node *Node) Search(no int) *Node {
	if node.Value.NO == no {
		return node
	} else if node.Value.NO > no && node.Left != nil {
		return node.Left.Search(no)
	} else if node.Value.NO < no && node.Right != nil {
		return node.Right.Search(no)
	}
	return nil
}

// 前序遍历打印：当前节点 --> 左节点 --> 右节点
func Print(current_node *Node) {
	fmt.Println(current_node.Value.NO)
	if current_node.Left != nil {
		Print(current_node.Left)
	}
	if current_node.Right != nil {
		Print(current_node.Right)
	}
}
