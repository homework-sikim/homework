package main

import "fmt"

type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

type BianryTree struct {
	Root *Node
}

func (b *BianryTree) Search(target int) bool {
	node := b.Root
	for {
		if node.Data == target {
			return true
		} else if node.Data > target {
			if node.Left != nil {
				node = node.Left
			} else {
				break
			}
		} else {
			if node.Right != nil {
				node = node.Right
			} else {
				break
			}
		}
	}
	return false
}

func main() {
	var n1, n2, n3, n4, n5, n6, n7 Node
	n1 = Node{Left: &n2, Data: 50, Right: &n3}
	n2 = Node{Left: &n4, Data: 25, Right: &n5}
	n3 = Node{Left: &n6, Data: 75, Right: &n7}
	n4 = Node{Data: 10}
	n5 = Node{Data: 33}
	n6 = Node{Data: 56}
	n7 = Node{Data: 89}

	tree := BianryTree{Root: &n1}
	fmt.Println(tree.Search(50))
	fmt.Println(tree.Search(25))
	fmt.Println(tree.Search(89))
	fmt.Println(tree.Search(51))
	fmt.Println(tree.Search(9))
}
