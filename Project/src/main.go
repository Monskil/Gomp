package main

import (
	"driver"
	//"fmt"
	"global"
	"network"
	"queue"
	//"flag"
	//"network/bcast"
	//"network/localip"
	//"network/peers"
	//"time"
	"fsm"
	"ordermanager"
)

/*
func test_button(){
	value := 0
	driver.Set_button_lamp(global.BUTTON_UP, global.FLOOR_2, global.ON)
	for {
		value = driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2)
		fmt.Println(value)
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2) != 0 {
			fmt.Println("Hello, you pressed button up floor 2 hehe.")
		}
	}*/

func main() {
	driver.Elevator_init()
	//obal var Global_order_list [global.NUM_GLOBAL_ORDERS]queue.Order

	//Global_order_list [global.NUM_GLOBAL_ORDERS]queue.Order

	new_order_chan := make(chan queue.Order, 10) //Er buffra til 10 for da får alle mulig ebestillinger plass
	updated_order_chan := make(chan queue.Order, 10)
	external_order_list_chan := make(chan [global.NUM_GLOBAL_ORDERS]queue.Order)
	internal_order_list_chan := make(chan [global.NUM_INTERNAL_ORDERS]queue.Order)
	new_order_bool_chan := make(chan bool)
	//isMaster := make(chan bool)
	//heQueue := make(chan [global.NUM_GLOBAL_ORDERS]queue.Order)
	//heQueue <- global_order_list

	go network.Network_info()
	go fsm.State_handler(new_order_bool_chan, updated_order_chan, external_order_list_chan, internal_order_list_chan)
	go queue.Handle_orders(new_order_bool_chan, new_order_chan, updated_order_chan, external_order_list_chan, internal_order_list_chan)
	go ordermanager.Detect_button_pressed(new_order_chan)
	//go ordermanager.HandleNewGlobal(external_order_list_chan, internal_order_list_chan)

	for { /*
			select {
			case temp := <-newButton:
				t := <-theQueue
				fmt.Println(temp)
				fmt.Println("lista er nå : ", t)

			}*/
	}

	// -- koble på nettverk

	// -- start ta imot bestillinger modus

	// testing
	//test_button()
	//queue.Init_queue()
	//network.Test_network()
	//ordermanager.Event_management()

}
