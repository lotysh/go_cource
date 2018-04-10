package main

import (
	"io"
	"strings"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	buf := make([]byte, 4)
	out := ""
	for {
		n, err := r.Read(buf)
		out += string(buf[:n])
		if err == io.EOF {
			break
		}
	}
	return out
}

func main() {
	const lines = `
abcd
234
dfadg`

	r := strings.NewReader(lines)
	text := ReadAll(r)
	println("Strings: ", text)
}