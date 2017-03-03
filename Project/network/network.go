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


type Global_order struct{
	text string
}

//type Slave_msg struct{
//	text string
//}

func TestNetwork(){
	//Setter opp for det master skal sende, og slavene motta
	global_order_sender := make(chan Global_order)
	global_order_receiver := make(chan Global_order)

	go bcast.Transmitter(30000, global_order_sender)
	go bcast.Receiver(30000, global_order_receiver)

	//Setter opp det slavene skal sende, og masteren motta
	/*slave_msg_sender := make(chan Slave_msg)
	slave_msg_receiver := make(chan Slave_msg)

	go bcast.Transmitter(30000, slave_msg_sender)
	go bcast.Receiver(30000, slave_msg_receiver)*/

	go func(){
		//S_msg := Slave_msg{"Slavemelding"} 
		G_msg := Global_order{"Global ordremelding"}
		for {
			//slave_msg_sender <- S_msg
			//time.Sleep(1*time.Second)
			global_order_sender <- G_msg
			time.Sleep(1*time.Second)
		}
	}()

	for {
		select{
		case a := <- global_order_receiver:
			fmt.Println("Received this a : ", a)
			time.Sleep(1*time.Second)
		//case b := <- slave_msg_receiver :
		//	fmt.Println("Received this b: ", b)
		}

	}
}
