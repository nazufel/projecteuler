package main

import (
	"fmt"
	"strconv"
)

/*
Question 2:
By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.
*/

// Function that takes n number of ints and returns ints
func FibonacciLoop(n int) int {
	// make a list named f, with two values, n+1 and n+2
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func main() {
	// control loop to calcuate the Fibonacci numbers up to i number of passes
	for i := 0; i <= 10; i++ {
		fmt.Print(strconv.Itoa(FibonacciLoop(i)) + " ")
	}
	fmt.Println("")
}
