package main

import "fmt"

/*
Question 1:
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func main() {
	number := [1000]int{}
	var total int
	for value := range number {
		if value%3 == 0 || value%5 == 0 {
			fmt.Println(value)
			total += value
		}
	}
	fmt.Println(total)
}
