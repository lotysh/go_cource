package main

import "fmt"

func main() {

    // Declare a variable called i which is a slice of 5 int.
	var i [5]int
	var f [9]float64
	//var s [4]string {c}
	s := []string{}
    	s = append(s, "one", "two", "three", "four")
    // Declare a variable called f which is a slice of 9 float64.
    // Declare a variable called s which is a slice of 4 string.

    fmt.Println(len(i), len(f), len(s))
}