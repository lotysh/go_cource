package main

import (
	"fmt"
)

// func main() {
// 	var i = 0
// 	for i = 1; i < 11; i++ {
// 		fmt.Println(i)
// 	}

// 	var j int = 10
// 	for ; j > 0; j-- {
// 		fmt.Println(j)
// 	}
// }

func main() {
	// var i = 0;
	i:=0
	for i = 1; i < 11; i++ {
		fmt.Println(i)
	}

	//var j int = 10
	j := func() int { return 10 }()

	//j:=10
	for ; j > 0; j-- {
		fmt.Println(j)
	}
}



