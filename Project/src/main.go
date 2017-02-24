package main

import (
	"driver"
	"fmt"
	def "global"
	//"time"
)

func main() {
	driver.Elevator_init()

	value := 0
	driver.Set_button_lamp(def.BUTTON_UP, def.FLOOR_2, def.ON)
	for {
		value = driver.Get_button_signal(def.BUTTON_UP, def.FLOOR_2)
		fmt.Println(value)
		if driver.Get_button_signal(def.BUTTON_UP, def.FLOOR_2) != 0 {
			fmt.Println("Hello, you pressed button up floor 2 hehe.")
		}
	}

	// -- koble på nettverk

	// -- start ta imot bestillinger modus
}
