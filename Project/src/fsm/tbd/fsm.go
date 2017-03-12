package fsm

import (
	"driver"
	"fmt"
	"global"
	"queue"
)

/* Hva skal fsm gjøre?
- Sjekke elev state og si hva den skal gjøre
*/

const (
	Idle int = iota
	Moving
	Door_open
	Stuck
)

func State_handler(updated_order_chan chan queue.Order, global_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order) {
	fmt.Println("Hello from state")
	elev_state := Idle
	current_order_chan := make(chan queue.Order)
	var current_order queue.Order
	for {
		switch elev_state {
		case Idle:
			current_order = event_idle(global_order_list_chan, internal_order_list_chan, current_order_chan)
			elev_state = Moving
		case Moving:
			//check if stuck -> Stuck (with timer, if u have been moving more than 10 sec)
			//check if reached floor (always turn on floor light when u r) -> check_if_order -> state = Door_open
			event_moving(updated_order_chan, current_order)
			elev_state = Door_open
		case Door_open:
			event_door_open(updated_order_chan, current_order)
			elev_state = Idle
		case Stuck:
			//call for help = tell master you cant move
			//check if your floor is changed, then you are not stuck anymore. Tell the master
		}
	}
}

func event_idle(global_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order, current_order_chan chan queue.Order) queue.Order {
	fmt.Println("Running event: Idle.")

	var current_order queue.Order
	// -- burde ta inn nåværende liste og ikke starte med en tom en
	var internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order
	var global_order_list [global.NUM_GLOBAL_ORDERS]queue.Order //hvis vi går inn i idle med full liste skjer det ingenting før ny ordre kommer

	for {
		for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
			if internal_order_list[i].Order_state != queue.Inactive {
				current_order = internal_order_list[i]
				return current_order
			}
		}
		for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
			if global_order_list[i].Order_state != queue.Inactive {
				current_order = internal_order_list[i]
				return current_order
			}
		}
		fmt.Println("going into select")
		select {
		case catch_internal_list := <-internal_order_list_chan:
			internal_order_list = catch_internal_list
		case catch_global_list := <-global_order_list_chan:
			global_order_list = catch_global_list
		}
	}
}

func event_moving(updated_order_chan chan queue.Order, current_order queue.Order) {
	fmt.Println("Running event: Moving.")

	current_order.Order_state = queue.Executing
	updated_order_chan <- current_order
	driver.Elevator_to_floor(current_order.Floor)

	// -- Her må det lages noe som gjør at den kan hente folk på veien
	// ---- dvs sjekke hver gang den kommer til en etasje om den har en bestilling der

	/*
		for{
			if driver.Get_floor_sensor_signal() != -1 {
				this_floor = driver.Get_floor_sensor_signal_floor_t()
				driver.Set_floor_indicator_lamp(this_floor)
				//if current_order.Floor == this_floor { //Endre dette til is there an order in this floor stuff
				//	break
				}


			}
	*/
}

func event_door_open(updated_order_chan chan queue.Order, current_order queue.Order) {
	fmt.Println("Running event: Door open.")
	driver.Open_door()
	current_order.Order_state = queue.Finished
	//fmt.Println(queue.Finished)
	updated_order_chan <- current_order
}
