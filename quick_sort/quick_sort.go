package main

import "fmt"

type Arr struct {
	data  []int
	left  int
	right int
	pivot int
}

func NewArr(array []int) Arr {
	return Arr{
		data:  array,
		left:  0,
		right: len(array) - 2,
		pivot: len(array) - 1,
	}
}

func main() {
	arr := NewArr([]int{44, 1, 4, 3, 7, 9, 12, 6, 2, 33, 8, 10, 5, 100, 21})
	arr.Split()
	fmt.Println(arr.data)
}

func (a *Arr) Split() {
	if len(a.data) < 2 {
		return
	}

	for {
		a.leftMove()
		a.rightMove()
		if a.left >= a.right {
			break
		}
		a.swap(a.left, a.right)
	}
	if a.data[a.left] > a.data[a.pivot] {
		a.swap(a.left, a.pivot)
	}

	leftArr := NewArr(a.data[0:a.left])
	leftArr.Split()

	rightArr := NewArr(a.data[a.left+1:])
	rightArr.Split()

}

func (a *Arr) leftMove() {
	for a.data[a.left] < a.data[a.pivot] {
		if a.left == len(a.data)-1 {
			break
		}
		a.left++
	}
}

func (a *Arr) rightMove() {
	for a.data[a.right] > a.data[a.pivot] {
		if a.right == 0 {
			break
		}
		a.right--
	}
}

func (a *Arr) swap(left, right int) {
	a.data[left], a.data[right] = a.data[right], a.data[left]
}
