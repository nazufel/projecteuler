package main

import "fmt"

/*
Question 2:
By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.
*/

func workingLoop(n int) []int {
	// make a slice of
	fmt.Println("n var passed to workingLoop is:", n)
	f := make([]int, 2)
	fmt.Println(f)
	/*if n < 2 {
		f = f[:]
	}
	*/
	f[0] = 0
	f[1] = 1
	for n <= 100 {
		f[n] = f[n-1] + f[n-2]
		n++
	}
	return f
}

func main() {
	n := 1
	fmt.Println(workingLoop(n))
}
