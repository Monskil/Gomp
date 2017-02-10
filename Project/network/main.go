package main

import (
	"./network/bcast"
	//"./network/localip"
	//"./network/peers"
	//"flag"
	"fmt"
	//"os"
	"time"
)

type HelloMsg struct {
	Message string
	Iter    int
}

func main() {

	//Make channels for sending and receiving HelloMsg
	hello_sender := make(chan HelloMsg)
	hello_receiver := make(chan HelloMsg)

	//Sier hvilken socket som skal gjøre hva
	go bcast.Transmitter(30000, hello_sender)
	go bcast.Receiver(30000, hello_receiver)

	go func() {
		message := HelloMsg{"hallopådo ", 0}
		for {
			message.Iter++
			hello_sender <- message
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case a := <-hello_receiver:
			fmt.Println("Receiving: ", a)
			time.Sleep(1 * time.Second)
		}
	}
}
