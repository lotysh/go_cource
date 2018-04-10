package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func countLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var linesc int
	for sc.Scan() {
		linesc++
	}
	return linesc, sc.Err()
}

func main() {
	linesc, err := countLines(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(linesc)
}