package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go coder()
	}

	for i := 0; i < 10; i++ {
		go tester()
	}

	go t.open()

	select {}
}

func coder() {
	work()
	t.coder <- struct{}{}
}

func tester() {
	work()
	t.tester <- struct{}{}
}

func work() {
	time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
}

func play() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

type table struct {
	coder  chan struct{}
	tester chan struct{}
}

func (t *table) open() {
	for {
		select {
		case <-t.coder:
			log.Println("coder waiting")
			<-t.tester
			log.Println("tester comes")
			log.Println("start pingpong")
			play()
			log.Println("end pingpong")
		case <-t.tester:
			log.Println("tester waiting")
			<-t.coder
			log.Println("coder comes")
			log.Println("start pingpong")
			play()
			log.Println("end pingpong")
		}
	}
}

var t = &table{
	coder:  make(chan struct{}),
	tester: make(chan struct{}),
}
