package _4_tree

import "fmt"

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

func (this *Node) PreOrder() {
	if this == nil {
		return
	}

	fmt.Println(this.data)
	this.left.PreOrder()
	this.right.PreOrder()
}

func (this *Node) InOrder() {
	if this == nil {
		return
	}

	this.left.InOrder()
	fmt.Println(this.data)
	this.right.InOrder()
}

func (this *Node) PostOrder() {
	if this == nil {
		return
	}

	this.left.PostOrder()
	this.right.PostOrder()
	fmt.Println(this.data)
}
