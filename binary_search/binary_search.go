package main

import (
	"fmt"
)

var array []int
var min int
var max int

func initArray() {
	for i := 0; i <= 6; i++ {
		array = append(array, i*2+1)
	}
	min = 0
	max = len(array)

}

func binarySearch(target int) {
	cnt := 0
	for {
		cnt++
		median := (max + min) / 2
		if array[median] > target {
			max = median - 1
			continue
		} else if array[median] < target {
			min = median + 1
			continue
		} else {
			fmt.Println("cnt : ", cnt)
			fmt.Println("index : ", median)
			break
		}
	}

}

func main() {
	initArray()
	binarySearch(13)
}
