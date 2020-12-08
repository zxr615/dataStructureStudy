package single

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Data: value}
}

func NewEmptyNode() *Node {
	return &Node{}
}

type LinkList struct {
	HeadNode *Node
	Length   int
}

func New() *LinkList {
	return &LinkList{}
}

func (list *LinkList) Add(value interface{}) *Node {
	node := NewNode(value)
	node.Next = list.HeadNode
	list.HeadNode = node
	list.Length++
	return node
}

func (list *LinkList) Append(value interface{}) *Node {
	if list.Empty() {
		return list.Add(value)
	}

	node := NewNode(value)
	currentNode := list.HeadNode
	for {
		// 找到最后一个节点
		if currentNode.Next == nil {
			currentNode.Next = node
			break
		}
		currentNode = currentNode.Next
	}
	list.Length++

	return node
}

func (list *LinkList) Insert(index int, value interface{}) *Node {
	if index <= 1 {
		return list.Add(value)
	}

	if index > list.Length {
		return list.Append(value)
	}

	currentNode := list.HeadNode
	node := NewNode(value)
	pre := NewEmptyNode()
	for k := 1; k <= index; k++ {
		if k == index-1 {
			pre = currentNode
		}
		if k == index {
			pre.Next = node
			node.Next = currentNode
			break
		}
		currentNode = currentNode.Next
	}

	list.Length++
	return node
}

func (list *LinkList) Remove(index int) bool {
	currentNode := list.HeadNode
	if index <= 1 {
		list.HeadNode = currentNode.Next
	}

	preNode := NewEmptyNode()
	for k := 1; k <= index; k++ {
		if k == index-1 {
			preNode = currentNode
		}

		if k == index {
			preNode.Next = currentNode.Next
		}

		currentNode = currentNode.Next
	}

	list.Length--
	return true
}

func (list *LinkList) Empty() bool {
	if list.Length == 0 {
		return true
	}
	return false
}

func (list *LinkList) Show() {
	currentNode := list.HeadNode
	for {
		fmt.Printf("%v\t", currentNode.Data)
		if currentNode.Next != nil {
			currentNode = currentNode.Next
		} else {
			break
		}
	}
}
