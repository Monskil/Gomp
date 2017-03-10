package ordermanager

import (
	"driver"
	//"fmt"
	"global"
	"queue"
	"time"
)

// -- we have to get the orders from the buttons
// -- send these on channels?
// -- check fsm
// -- try making an "easy" version for the internal orders first

func Detect_button_pressed(newButton chan queue.Order) {
	var order queue.Order

	for {

		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_1) != 0 {
			order.Button = global.BUTTON_UP
			order.Floor = global.FLOOR_1
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2) != 0 {
			order.Button = global.BUTTON_UP
			order.Floor = global.FLOOR_2
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order

			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_DOWN, global.FLOOR_2) != 0 {
			order.Button = global.BUTTON_DOWN
			order.Floor = global.FLOOR_2
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_3) != 0 {
			order.Button = global.BUTTON_UP
			order.Floor = global.FLOOR_3
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_DOWN, global.FLOOR_3) != 0 {
			order.Button = global.BUTTON_DOWN
			order.Floor = global.FLOOR_3
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_DOWN, global.FLOOR_4) != 0 {
			order.Button = global.BUTTON_DOWN
			order.Floor = global.FLOOR_4
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_1) != 0 {
			order.Button = global.BUTTON_COMMAND
			order.Floor = global.FLOOR_1
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_2) != 0 {
			order.Button = global.BUTTON_COMMAND
			order.Floor = global.FLOOR_2
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_3) != 0 {
			order.Button = global.BUTTON_COMMAND
			order.Floor = global.FLOOR_3
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_4) != 0 {
			order.Button = global.BUTTON_COMMAND
			order.Floor = global.FLOOR_4
			order.Order_state = queue.Active
			order.Assigned_to = global.ELEV_1
			newButton <- order
			time.Sleep(1 * time.Second)

		}

	}
}

/*func detect_internal_button_pressed() (global.Button_t, bool) {
	if driver.Get_button_signal(global.BUTTON_COMMAND1) {
		return global.BUTTON_COMMAND1, true
	}
	if driver.Get_button_signal(global.BUTTON_COMMAND2) {
		return global.BUTTON_COMMAND2, true
	}
	if driver.Get_button_signal(global.BUTTON_COMMAND3) {
		return global.BUTTON_COMMAND3, true
	}
	if driver.Get_button_signal(global.BUTTON_COMMAND1) {
		return global.BUTTON_COMMAND4, true
	} else {
		break
	}
}*/
/*
func Event_management() {
	/*value := 0
	driver.Set_button_lamp(global.BUTTON_UP, global.FLOOR_2, global.ON)
	for {
		value = driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2)
		fmt.Println(value)
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2) != 0 {
			fmt.Println("Hello, you pressed button up floor 2 hehe.")
		}*/ /*
	var global_order_list [global.NUM_GLOBAL_ORDERS]queue.Order

	i := 1
	for i < 3 {
		b, f, t := Detect_external_button_pressed()
		//fmt.Println(b, f, t)
		if t {
			var order1 = queue.Make_new_order(b, f, queue.Finished, global.ELEV_1)
			queue.Add_new_global_order(order1, global_order_list)
			fmt.Println(b, f, t)

		}
		time.Sleep(1 * time.Second)

	}

	// if button is pressed
	// -- what should happen?
}*/

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
