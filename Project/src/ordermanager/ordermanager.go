package ordermanager

import (
	"driver"
	"fmt"
	"global"
	"queue"
	"time"
)

//-- We must make a function for this..

func Make_new_order(button global.Button_t, floor global.Floor_t, state global.Order_state, elev global.Assigned_to)queue.Order{
	var order queue.Order
	order.Button = button
	order.Floor = floor
	order.Order_state = state
	order.Assigned_to = elev //Her bør kanskje kost-funksjonen kjøres?? Iallefall hvis heisens state er Inactive
	return order
	//PÅ ALLE COMMAND, ASSIGNED TO BØR VÆRE SEG SELV.
}

func Detect_button_pressed(new_order_chan chan queue.Order) {
	fmt.Println("Running: Detect button pressed. ")
	var order queue.Order

	for {

		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_1) != 0 {
			order = queue.Make_new_order(global.BUTTON_UP, global.FLOOR_1, queue.Active, global.NONE)
			new_order_chan <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_2) != 0 {
			order = queue.Make_new_order(global.BUTTON_UP, global.FLOOR_2, queue.Active, global.NONE)
			new_order_chan <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_DOWN, global.FLOOR_2) != 0 {
			order = queue.Make_new_order(global.BUTTON_DOWN, global.FLOOR_2, queue.Active, global.NONE)
			new_order_chan <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_UP, global.FLOOR_3) != 0 {
			order = queue.Make_new_order(global.BUTTON_UP, global.FLOOR_3, queue.Active, global.NONE)
			new_order_chan <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_DOWN, global.FLOOR_3) != 0 {
			order = queue.Make_new_order(global.BUTTON_DOWN, global.FLOOR_3, queue.Active, global.NONE)
			new_order_chan <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_DOWN, global.FLOOR_4) != 0 {
			order = queue.Make_new_order(global.BUTTON_DOWN, global.FLOOR_4, queue.Active, global.NONE)
			new_order_chan <- order
			time.Sleep(1 * time.Second)
		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_1) != 0 {
			order = queue.Make_new_order(global.BUTTON_COMMAND, global.FLOOR_1, queue.Active, global.ELEV_1)
			new_order_chan <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_2) != 0 {
			order = queue.Make_new_order(global.BUTTON_COMMAND, global.FLOOR_2, queue.Active, global.ELEV_1)
			new_order_chan <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_3) != 0 {
			order = queue.Make_new_order(global.BUTTON_COMMAND, global.FLOOR_3, queue.Active, global.ELEV_1)
			new_order_chan <- order
			time.Sleep(1 * time.Second)

		}
		if driver.Get_button_signal(global.BUTTON_COMMAND, global.FLOOR_4) != 0 {
			order = queue.Make_new_order(global.BUTTON_COMMAND, global.FLOOR_4, queue.Active, global.ELEV_1)
			new_order_chan <- order
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
