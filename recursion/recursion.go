package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func fibonacci(n int) int {
	if n < 3 {
		return 1
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

var hanoiCnt = 0

func hanoiCnTAdd() {
	hanoiCnt = hanoiCnt + 1
}

func hanoiMove(n int, a int, b int, c int) {
	hanoiCnTAdd()
	if n == 1 {
		fmt.Println(n, ",", a, "->", c)
		return
	}
	hanoiMove(n-1, a, b, c)
	fmt.Println(n, ",", a, "->", c)
	hanoiMove(n-1, b, c, a)
}

func main() {
	fmt.Println(factorial(6))
	fmt.Println(fibonacci(6))

	hanoiMove(3, 1, 2, 3)
	fmt.Println(hanoiCnt)
}
