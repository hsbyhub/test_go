package test

import (
	"fmt"
	"testing"
)

var end chan int

func read(ch <-chan int) {
	fmt.Println("before read")
	fmt.Println("read:", <-ch)
	fmt.Println("after read")
	end <- 1
}

func write(ch chan<- int) {
	fmt.Println("before write")
	ch <- 100
	fmt.Println("after write")
}

func TestChannel(t *testing.T) {
	end = make(chan int)
	ch := make(chan int, 1)
	go write(ch)
	go read(ch)
	<-end
}
