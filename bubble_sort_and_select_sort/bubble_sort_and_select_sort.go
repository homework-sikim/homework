package main

import "fmt"

func initArray(length int) []int {
	a := []int{}
	for i := length; i > 0; i-- {
		a = append(a, i)
	}
	return a
}

func bubbleSort(a []int) []int {
	cnt := 0
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			cnt++
			if a[j] > a[j+1] {
				cnt++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(cnt)
	return a
}

func selectSort(a []int) []int {
	cnt := 0
	for i := 0; i < len(a); i++ {
		minIndex := i
		// i와 minIndex 는 같으므로 j가 i부터 시작하느 이유가 없습니다.
		for j := i + 1; j < len(a); j++ {
			cnt++
			if a[j] <= a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	fmt.Println(cnt)
	return a
}

func main() {
	a := initArray(20)
	bubbleSort(a)
	fmt.Println(a)

	a = initArray(20)
	selectSort(a)
	fmt.Println(a)
}
