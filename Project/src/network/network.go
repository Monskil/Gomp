package network

import (
	"bcast"
	"fmt"
	"time"
)

// -- kan bruke json marshal greier for å pakke meldingen, og unpakke den
// -- sender det da som bytes, må være en public struct (stor forbokstav)


type Master_msg struct {
	Text string
}

type Slave_msg struct{
	Text string
}

func Test_network() {

	//Make channels for sending and receiving HelloMsg
	master_sender := make(chan Master_msg)
	master_receiver := make(chan Master_msg)
	slave_sender := make(chan Slave_msg)
	slave_receiver := make(chan Slave_msg)

	//Sier hvilken socket som skal gjøre hva
	go bcast.Transmitter(30000, master_sender)
	go bcast.Receiver(30000, master_receiver)
	go bcast.Transmitter(30000, slave_sender)
	go bcast.Receiver(30000, slave_receiver)

	go func() {
		master_message := Master_msg{"Hellooo, master message."}
		slave_message := Slave_msg{"Hellooo, slave message."}
		for {
			master_sender <- master_message
			slave_sender <- slave_message
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case master := <- master_receiver:
			fmt.Println("Receiving: ", master)
			time.Sleep(1 * time.Second)
		case slave := <- slave_receiver:
			fmt.Println("Receiving: ", slave)
			time.Sleep(1 * time.Second)
		}
	}
}
