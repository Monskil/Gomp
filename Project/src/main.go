package main

import (
	"driver"
	//fmt"
	//global"
	"queue"
	//"network"
	//"flag"
	//"network/bcast"
	//"network/localip"
	//"network/peers"
	//"time"
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

	newButton := make(chan queue.Order, 10) //Er buffra til 10 for da får alle mulig ebestillinger plass
	//heQueue := make(chan [global.NUM_GLOBAL_ORDERS]queue.Order)
	//heQueue <- global_order_list

	go ordermanager.Detect_external_button_pressed(newButton)
	go queue.Handle_button_pressed(newButton)

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
