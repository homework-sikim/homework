package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	First  *Node
	Last   *Node
	Length int
}

type Node struct {
	Prev *Node
	Data int
	Next *Node
}

func (l *LinkedList) Push(data int) {
	node := Node{Prev: nil, Data: data, Next: nil}
	if l.First == nil {
		l.First = &node
		l.Last = &node
		l.Length += 1
	} else {
		node.Next = l.First
		l.First.Prev = &node
		l.First = &node
		l.Length += 1
	}
}

func (l *LinkedList) Get(index int) (int, error) {
	if index < 0 || index > l.Length {
		return 0, errors.New("index out of range")
	} else {
		temp := l.First
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
		return temp.Data, nil
	}
}

func (l *LinkedList) Find(target int) bool {
	if l.Length == 0 {
		return false
	} else {
		temp := l.First
		for i := 0; i < l.Length; i++ {
			if temp.Data == target {
				return true
			}
			temp = temp.Next
		}
		return false
	}
}

func (l *LinkedList) Delete() bool {
	if l.Length == 0 {
		return false
	}

	temp := l.First.Next
	if temp != nil {
		temp.Prev = nil
		l.First = temp
		l.Length -= 1
	} else {
		l.First = nil
		l.Last = nil
		l.Length = 0
	}
	return true

}

func main() {
	linkedList := LinkedList{}
	linkedList.Push(1)
	linkedList.Push(2)
	linkedList.Push(3)
	linkedList.Push(4)
	data, err := linkedList.Get(1)
	fmt.Println(data, err)
	data, err = linkedList.Get(5)
	fmt.Println(data, err)
	fmt.Println(linkedList.Find(1))
	fmt.Println(linkedList.Find(5))

	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
	linkedList.Push(1)
	fmt.Println(linkedList)
	linkedList.Push(2)
	fmt.Println(linkedList)
	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
}
