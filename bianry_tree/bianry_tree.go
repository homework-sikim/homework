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
	Status int // 0: root, 1:left , 2:right
}

type BianryTree struct {
	Root *Node
}

func (b *BianryTree) SearchNode(target int) *Node {
	node := b.Root
	for {
		if node.Data == target {
			return node
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
	return nil
}

func (b *BianryTree) SearchEmptyNode(target int) *Node {
	node := b.Root
	for {
		if node.Data == target {
			return nil
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
	return node
}

func (b *BianryTree) Insert(newVal int) error {
	node := Node{}
	node.Data = newVal

	if b.Root == nil {
		b.Root = &node
	} else {
		result := b.SearchEmptyNode(newVal)
		if result == nil {
			return errors.New("값이 존재한다.")
		}
		node.Parent = result
		if result.Data > newVal {
			node.Status = 1
			result.Left = &node
		} else {
			node.Status = 2
			result.Right = &node
		}
	}
	return nil
}

func SearchLeftEmptyNode(targetNode *Node) *Node {
	node := targetNode
	for {
		if node.Left != nil {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

func (b *BianryTree) InsertRightChild(node *Node) error {
	for {
		err := b.Insert(node.Data)
		if err != nil {
			return err
		}
		if node.Left != nil || node.Right != nil {
			if node.Left != nil {
				err = b.InsertRightChild(node.Left)
				if err != nil {
					return err
				}
			}

			if node.Right != nil {
				err = b.InsertRightChild(node.Right)
				if err != nil {
					return err
				}
			}

		} else {
			return nil
		}
	}
	return nil
}

func (b *BianryTree) Delete(target int) error {
	// 1. 삭제 타겟을 찾는다.
	node := b.SearchNode(target)
	if node == nil {
		return errors.New("값이 없다.")
	}

	// 2. 자식이 없다면 그냥 삭제한다.
	if node.Left == nil && node.Right == nil {
		if node.Status == 1 {
			node.Parent.Left = nil
			node = nil
			return nil
		} else if node.Status == 2 {
			node.Parent.Right = nil
			node = nil
			return nil
		} else {
			b.Root = nil
		}
	} else if node.Left == nil || node.Right == nil {
		// 3. 타겟의 자식이 하나면 자식을 삭제된 노드 위치에 넣는다 .
		if node.Left != nil {
			temp := node.Left
			temp.Parent = node.Parent
			temp.Status = node.Status
			node = temp
			node.Parent.Left = node
		} else {
			temp := node.Right
			temp.Parent = node.Parent
			temp.Status = node.Status
			node = temp
			node.Parent.Right = node
		}

		if node.Status == 0 {
			b.Root = node
		}
	} else {
		// 4. 후속자 노드를 찾는다.
		temp := SearchLeftEmptyNode(node.Right)
		tempRight := temp.Right

		//5. 삭제타겟자리에 후속자 노드를 넣는다.
		temp.Parent = node.Parent
		temp.Status = node.Status
		temp.Right = nil
		temp.Left = node.Left
		node = temp
		node.Left.Parent = node

		if node.Status == 0 {
			b.Root = node
		} else if node.Status == 1 {
			node.Parent.Left = node
		} else {
			node.Parent.Right = node
		}

		//6. 남은 오른쪽 자식은 추가한다.
		if tempRight != nil {
			err := b.InsertRightChild(tempRight)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func main() {
	tree := BianryTree{}
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(10)
	tree.Insert(35)
	tree.Insert(56)
	tree.Insert(89)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(52)
	tree.Insert(82)
	tree.Insert(95)
	tree.Insert(32)
	tree.Insert(31)
	tree.Insert(45)
	tree.Insert(43)

	fmt.Println(tree.SearchNode(35))
	fmt.Println(tree.SearchNode(40))
	fmt.Println(tree.SearchNode(30))
	fmt.Println(tree.SearchNode(32))
	fmt.Println(tree.SearchNode(31))
	fmt.Println(tree.SearchNode(45))
	fmt.Println(tree.SearchNode(43))
	tree.Delete(35)
	fmt.Println(tree.SearchNode(25))
	fmt.Println(tree.SearchNode(40))
	fmt.Println(tree.SearchNode(30))
	fmt.Println(tree.SearchNode(32))
	fmt.Println(tree.SearchNode(31))
	fmt.Println(tree.SearchNode(45))
	fmt.Println(tree.SearchNode(43))
}
