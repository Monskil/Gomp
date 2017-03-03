package network

import (
	"./bcast"
	//"./network/localip"
	//"./network/peers"
	//"flag"
	"fmt"
	//"os"
	"time"
)

type SlaveMsg struct {
	Message string
}

type GlobalMsg struct{
	text string
}

func Netrun() {

	//Make channels for sending and receiving HelloMsg
	slave_sender := make(chan SlaveMsg)
	slave_receiver := make(chan SlaveMsg)

	//Sier hvilken socket som skal gjøre hva
	go bcast.Transmitter(30000, slave_sender)
	go bcast.Receiver(30000, slave_receiver)


	global_sender := make(chan GlobalMsg)
	global_receiver := make(chan GlobalMsg)

	go bcast.Transmitter(30000, global_sender)
	go bcast.Receiver(30000, global_receiver)

	go func() {
		message := SlaveMsg{"Slavesender"}
		globmessage := GlobalMsg{"Globalsender"}
		for {
			slave_sender <- message
			time.Sleep(1*time.Second)
			global_sender <- globmessage
			time.Sleep(1*time.Second)
		}
	}()

	for {
		select {
		case a := <-slave_receiver:
			fmt.Println("Receiving: ", a)
		case b := <- global_receiver:
			fmt.Println("Receiving :" , b)
		}
	}
}
