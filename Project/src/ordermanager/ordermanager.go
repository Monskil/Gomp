package order_manager

import (
	"driver"
	"fmt"
)

// -- we have to get the orders from the buttons
// -- send these on channels?
// -- check fsm
// -- try making an "easy" version for the internal orders first

func detect_button_pressed(){
	driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2)
}

func Order_management() {
	// if button is pressed
	// -- what should happen?
}
