package network

import (
	"../queue"
	"bcast"
	"fmt"
	"global"
	"time"
)

// -- kan bruke json marshal greier for å pakke meldingen, og unpakke den
// -- sender det da som bytes, må være en public struct (stor forbokstav)

type Master_msg struct {
	Master_order_list [6]queue.Order
}

type Slave_msg struct {
	Slave_order_list [10]queue.Order
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

	// make lists
	var internal_order_list [NUM_INTERNAL_ORDERS]Order
	var global_order_list [NUM_GLOBAL_ORDERS]Order
	var my_order_list [NUM_ORDERS]Order

	// test example making orders
	var order1 = queue.Make_new_order(BUTTON_UP, FLOOR_2, finished, ELEV_2)
	var order2 = queue.Make_new_order(BUTTON_COMMAND, FLOOR_1, active, ELEV_3)
	var order3 = queue.Make_new_order(BUTTON_UP, FLOOR_1, active, ELEV_1)
	var order4 = queue.Make_new_order(BUTTON_COMMAND, FLOOR_2, active, ELEV_2)

	internal_order_list = queue.Add_new_internal_order(order1, internal_order_list)
	internal_order_list = queue.Add_new_internal_order(order2, internal_order_list)
	internal_order_list = queue.Add_new_internal_order(order2, internal_order_list)

	global_order_list = queue.Add_new_global_order(order2, global_order_list)
	global_order_list = queue.Add_new_global_order(order1, global_order_list)
	global_order_list = queue.Add_new_global_order(order2, global_order_list)
	global_order_list = queue.Add_new_global_order(order3, global_order_list)
	global_order_list = queue.Add_new_global_order(order4, global_order_list)
	my_order_list = queue.Make_my_order_list(internal_order_list, global_order_list)

	go func() {
		master_message := Master_msg{global_order_list}
		slave_message := Slave_msg{my_order_list}
		for {
			master_sender <- master_message
			slave_sender <- slave_message
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case master := <-master_receiver:
			fmt.Println("Receiving: ", master)
			time.Sleep(1 * time.Second)
		case slave := <-slave_receiver:
			fmt.Println("Receiving: ", slave)
			time.Sleep(1 * time.Second)
		}
	}
}
