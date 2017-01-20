package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int
var channel = make(chan int, 1)

func createThread1() {
	channel <- 0
	for n := 0; n < 1000001; n++ {
		i++
	}
	<- channel
	
}

func createThread2() {
	channel <- 0
	for n := 0; n < 1000000; n++ {
		i--
	}
	<- channel
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go createThread1()
	go createThread2()

	time.Sleep(100 * time.Millisecond)
	Println(i)

}

