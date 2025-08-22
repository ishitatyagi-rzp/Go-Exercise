package main

import "fmt"

type Node struct {
	value string
	left  *Node
	right *Node
}

func PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	PreOrder(n.left)
	PreOrder(n.right)
}

func PostOrder(n *Node) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	fmt.Print(n.value, " ")
}
func main() {
	a := &Node{value: "a"}
	b := &Node{value: "b"}
	c := &Node{value: "c"}

	minus := &Node{value: "-", left: b, right: c}
	plus := &Node{value: "+", left: a, right: minus}

	fmt.Print("PreOrder Traversal:")
	PreOrder(plus)
	fmt.Println()

	fmt.Print("PostOrder Traversal:")
	PostOrder(plus)
	fmt.Println()
}
