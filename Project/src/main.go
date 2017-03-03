package main

import (
	"driver"
	"fmt"
	"global"
	//"queue"
	//"network"
	//"time"
)

func test_button(){
	value := 0
	driver.Set_button_lamp(global.BUTTON_UP, global.FLOOR_2, global.ON)
	for {
		value = driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2)
		fmt.Println(value)
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2) != 0 {
			fmt.Println("Hello, you pressed button up floor 2 hehe.")
		}
	}

func main() {
	driver.Elevator_init()

	// -- koble på nettverk

	// -- start ta imot bestillinger modus
	
	// testing
	test_button()
  	//queue.Init_queue()
	//network.Test_network()
}
