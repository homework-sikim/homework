package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Parent *Node
	Left   *Node
	Data   int
	Right  *Node
}

type BianryTree struct {
	Root *Node
}

func (b *BianryTree) Search(target int) (*Node, bool) {
	node := b.Root
	isLeft := true
	for {
		if node.Data == target {
			return node, isLeft
		} else if node.Data > target {
			if node.Left != nil {
				node = node.Left
			} else {
				break
			}
		} else {
			if node.Right != nil {
				isLeft = false
				node = node.Right
			} else {
				break
			}
		}
	}
	return nil, isLeft
}

func (b *BianryTree) Delete(target int) error {
	node, isLeft := b.Search(target)
	if node == nil {
		return errors.New("값이 없다.")
	}

	if node.Left == nil && node.Right == nil {
		if isLeft {
			node.Parent.Left = nil
			node = nil
			return nil
		} else {
			node.Parent.Right = nil
			node = nil
			return nil
		}
	}
	return nil
}

func main() {
	var n1, n2, n3, n4, n5, n6, n7 Node
	n1 = Node{Parent: nil, Left: &n2, Data: 50, Right: &n3}
	n2 = Node{Parent: &n1, Left: &n4, Data: 25, Right: &n5}
	n3 = Node{Parent: &n1, Left: &n6, Data: 75, Right: &n7}
	n4 = Node{Parent: &n2, Data: 10}
	n5 = Node{Parent: &n2, Data: 33}
	n6 = Node{Parent: &n3, Data: 56}
	n7 = Node{Parent: &n3, Data: 89}

	tree := BianryTree{Root: &n1}
	fmt.Println(tree.Delete(50))
	fmt.Println(tree.Search(25))
	fmt.Println(tree.Search(89))
	fmt.Println(tree.Search(51))
	fmt.Println(tree.Delete(9))
}
