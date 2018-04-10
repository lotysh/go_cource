package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// This program has two problems.
	// The first problem is the output reports the 1st prime
	// as the 0 th prime.
	// The second problem is the suffix the 1st prime should be
	// 1 st, not 1 th.
	// Can you fix both of these problems.

	for i, p := range primes {
		if i==0 {
		fmt.Println("The", i+1, "st prime is", p)
		} else 	if i==1 {
		fmt.Println("The", i+1, "st prime is", p)
		} else if i>1 {
		fmt.Println("The", i+1, "th prime is", p)
		}
	}
}


