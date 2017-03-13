package fsm

import (
	"driver"
	"fmt"
	"global"
	"queue"
	"time"
)

// Elevator states
const (
	Idle int = iota
	Moving
	Door_open
	Stuck
)

func State_handler(new_order_bool_chan chan bool, updated_order_chan chan queue.Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order) {
	fmt.Println("Running: State handler. ")
	elev_state := Idle
	var internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order
	var external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order
	//current_order_chan := make(chan queue.Order)
	var current_order queue.Order
	for {
		switch elev_state {
		case Idle:
			current_order, internal_order_list, external_order_list = event_idle(new_order_bool_chan, internal_order_list_chan, external_order_list_chan)
			elev_state = Moving
		case Moving:
			event_moving(current_order, updated_order_chan, internal_order_list, internal_order_list_chan, external_order_list, external_order_list_chan)
			elev_state = Door_open
		case Door_open:
			event_door_open(updated_order_chan, current_order)
			elev_state = Idle
		case Stuck:
			event_stuck()
			elev_state = Idle
		}
	}
}

func event_idle(new_order_bool_chan chan bool, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order) (queue.Order, [global.NUM_INTERNAL_ORDERS]queue.Order, [global.NUM_GLOBAL_ORDERS]queue.Order) {
	fmt.Println("Running event: Idle.")

	var current_order queue.Order
	//-- må endre til å ta inn nåværende liste og ikke starte med en tom en
	var internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order
	var external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order

	// Check if there is an order in one of the lists
	for {
		fmt.Println("At the top")
		for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
			if internal_order_list[i].Order_state != queue.Inactive {
				current_order = internal_order_list[i]
				return current_order, internal_order_list, external_order_list
			}
		}
		for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
			if external_order_list[i].Order_state != queue.Inactive {
				current_order = external_order_list[i]
				return current_order, internal_order_list, external_order_list
			}
		}

		// Wait until some order is added
		select {
		case catch_new_order := <-new_order_bool_chan:
			fmt.Println("Got new order bool ", catch_new_order, " in Idle.")
			//-- using catch_var to be able to use the lists outside of select
			catch_internal_order_list := <-internal_order_list_chan
			catch_external_order_list := <-external_order_list_chan
			internal_order_list = catch_internal_order_list
			external_order_list = catch_external_order_list

			fmt.Println("1", internal_order_list, external_order_list)

			// Push lists back to channels
			go queue.Internal_order_list_to_channel(internal_order_list, internal_order_list_chan)
			go queue.External_order_list_to_channel(external_order_list, external_order_list_chan)
			break
		}
	}
}

func event_moving(current_order queue.Order, updated_order_chan chan queue.Order, internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order, external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order) {
	fmt.Println("Running event: Moving.")

	// Go to state Stuck if the elevator is in state Moving for more than 12 seconds
	//-- Must implement a way to check if timeout when inside elevator to floor
	/*timer := time.NewTimer(12 * time.Second)
	timeout := false
	go func() {
		<-timer.C
		timeout = true
	}()*/

	var order_list [global.NUM_ORDERS]queue.Order

	order_list = queue.Make_my_order_list(internal_order_list, external_order_list)
	fmt.Println("My order list is now: ", order_list)

	// Set order state to executing
	current_order.Order_state = queue.Executing
	updated_order_chan <- current_order

	// Go to floor
	floor := current_order.Floor
	elevator_to_floor(floor, order_list, updated_order_chan, current_order)
}

func event_door_open(updated_order_chan chan queue.Order, current_order queue.Order) {
	fmt.Println("Running event: Door open.")

	// Open door
	driver.Open_door()
	driver.Set_button_lamp(current_order.Button, current_order.Floor, global.OFF) //-- can be moved to before open door

	// Set order state to finished
	current_order.Order_state = queue.Finished
	updated_order_chan <- current_order
}

func event_stuck() {
	// If the floor is changed, elevator is no longer stuck
	// Go to the first floor and set state to idle
	fmt.Println("Running event: Stuck.")
	my_floor := driver.Get_floor_sensor_signal()
	for {
		//-- Tell master the elevator is stuck (?)
		if my_floor != driver.Get_floor_sensor_signal() {
			driver.Elevator_to_floor_direct(global.FLOOR_1)
			break
		}
	}
}

func elevator_to_floor(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order, updated_order_chan chan queue.Order, current_order queue.Order) {
	current_floor_int := driver.Get_floor_sensor_signal()
	current_floor := driver.Floor_int_to_floor_t(current_floor_int)

	floor_int := driver.Floor_t_to_floor_int(floor)

	// Check if the elevator is between two floors
	timer := time.NewTimer(3 * time.Second)
	timeout := false
	go func() {
		<-timer.C
		timeout = true
	}()
	for driver.Get_floor_sensor_signal() == -1 {
		if !timeout {
			driver.Set_motor_direction(global.DIR_UP)
		} else if timeout {
			driver.Set_motor_direction(global.DIR_DOWN)
		}
	}

	// Go to desired floor
	fmt.Println(current_floor_int, floor_int)
	if current_floor_int < floor_int {
		fmt.Println("Going up.")
		driver.Set_motor_direction(global.DIR_UP)
		time.Sleep(100 * time.Millisecond)
		for driver.Get_floor_sensor_signal() != floor_int {
			current_floor = driver.Floor_int_to_floor_t(driver.Get_floor_sensor_signal())

			// When arriving at any floor, check for order
			if driver.Get_floor_sensor_signal() != -1 {
				pick_up_order_on_the_way(current_floor, order_list, updated_order_chan, current_order)
				time.Sleep(10 * time.Millisecond)
			}
		}

	} else if current_floor_int > floor_int {
		fmt.Println("Going down.")
		driver.Set_motor_direction(global.DIR_DOWN)

		for driver.Get_floor_sensor_signal() != floor_int {
			current_floor = driver.Floor_int_to_floor_t(driver.Get_floor_sensor_signal())

			// When we arrive at any floor, check for order
			if driver.Get_floor_sensor_signal() != -1 {
				pick_up_order_on_the_way(current_floor, order_list, updated_order_chan, current_order)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
	// Stop when at desired floor
	driver.Set_motor_direction(global.DIR_STOP)
}

func pick_up_order_on_the_way(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order, updated_order_chan chan queue.Order, current_order queue.Order) {
	// If the elevator has an order in this floor, stop and take this order
	order_in_floor := check_if_order_in_floor(floor, order_list)
	if order_in_floor {
		driver.Set_motor_direction(global.DIR_STOP)
		event_door_open(updated_order_chan, current_order)
	}
}

func check_if_order_in_floor(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order) bool {
	for i := 0; i < global.NUM_ORDERS; i++ {
		if order_list[i].Floor == floor && order_list[i].Order_state != queue.Inactive {
			return true
		}
	}
	return false
}
