package main

import "fmt"

func main() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	var z,j,i int
	y := len(primes)
	for i := 0; i < y; i++ {
		if primes[i] < 10 {
			fmt.Println(primes[i])
			z=i+1
			}
}	
	sub := primes[z:]
	fmt.Println(sub)
	
	// Write a program to print only the primes less than 10
	// loop through the slice of primes and test if the value
	// is less than 10. When you find a value that is 10 or more
	// slice the list of primes at that point and print it.

	// Bonus: write a print only the two digit primes.
	for j = z; j < y; j++ {
		if primes[j] > 99 {
			break
		}
	}
	fmt.Println(primes[i:j])


	// see src/slices/slices9/main.go for the answer
}
