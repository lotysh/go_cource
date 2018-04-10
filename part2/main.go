// Solution to part 2 of the Whispering Gophers code lab.
//
// This program extends part 1.
//
// It makes a connection the host and port specified by the -dial flag, reads
// lines from standard input and writes JSON-encoded messages to the network
// connection.
//
// You can test this program by installing and running the dump program:
// 	$ go get github.com/campoy/whispering-gophers/util/dump
// 	$ dump -listen=localhost:8000
// And in another terminal session, run this program:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "127.0.0.1:8001", "host:port to dial")
//var dialAddr = flag.String("dial", "192.168.56.1:63887", "host:port to dial")


type Message struct {
	Body string
}

func main() {
	// TODO: Parse the flags.
	flag.Parse()
	c, err := net.Dial("tcp", *dialAddr)
	//var dialAddr = flag.String("listen", "127.0.0.1:4000", "host:port to listen")
	// TODO: Open a new connection using the value of the "dial" flag.
	// TODO: Don't forget to check the error.
	if err != nil {
			log.Fatal(err)
		}
	s := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(c)

	// TODO: Create a json.Encoder writing into the connection you created before.
	for s.Scan() {
		m := Message{Body: s.Text()}
		err := enc.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
