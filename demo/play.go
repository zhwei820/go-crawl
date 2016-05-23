package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

var tokenCh chan string
var popchan chan string

func main() {
	StartTokenGen()
}

func StartTokenGen() {
	tokenCh = make(chan string)
	popchan = make(chan string)

	go TokenGen(tokenCh)
	//go TokenGen1(popchan)

	for {
		select {
		case t := <-tokenCh:
			log.Println(t)
		case p := <-popchan:
			log.Println(p)
		default:
			log.Println("no message")
		}
	}

}

func TokenGen1(ch chan string) {
	for {
		popchan <- "pop"
		time.Sleep(1 * time.Second)

	}
}

func TokenGen(ch chan string) {
	for {
		log.Println("goroutine")
		time.Sleep(1 * time.Second)

	}
}
func RegexpDemo() {
	r, _ := regexp.Compile(`/fb://page/(\d+)`)

	matcher := r.FindStringSubmatch("dsadfsdvdfb://page/123253412")

	fmt.Println(matcher)
}
