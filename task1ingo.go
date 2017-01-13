package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int

func createThread1() {
	for n := 0; n < 1000000; n++ {
		i++
	}
}

func createThread2() {
	for n := 0; n < 1000000; n++ {
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go createThread1()
	go createThread2()

	time.Sleep(100 * time.Millisecond)
	Println(i)

}
