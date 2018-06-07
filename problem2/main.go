package main

import "fmt"

// Function using recursion to calculate Fibonacci numbers
func fibonacciLoop(n int) int {
	for n <= 1 {
		return n
	}
	return fibonacciLoop(n-1) + fibonacciLoop(n-2)
}

func main() {
	total := 0
	for n := 0; n <= 33; n++ {
		if fibonacciLoop(n)%2 == 0 {
			fmt.Println(fibonacciLoop(n))
			total += fibonacciLoop(n)
			fmt.Println(total)
		}
	}
}
