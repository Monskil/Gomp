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

/*func Elev_state_to_chan(elev_state int, elev_state_chan chan int){
	elev_state_chan <- elev_state
}*/

func State_handler(new_order_bool_chan chan bool, updated_order_chan chan queue.Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order) {
	fmt.Println("Running: State handler. ")
	elev_state := Idle
	//elev_state_chan := make(chan int)
	var internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order
	var external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order

	//Elev_state_to_chan(Idle, elev_state_chan)

	/*door_timeout_chan := make(chan bool)
	door_timer_reset_chan := make(chan bool)

	go door_timer(door_timeout_chan, door_timer_reset_chan)*/

	var current_order queue.Order
	for {
		switch elev_state {
		case Idle:
			time.Sleep(100 * time.Millisecond)
			current_order, internal_order_list, external_order_list = event_idle(new_order_bool_chan, internal_order_list_chan, external_order_list_chan)
			elev_state = Moving
		case Moving:
			event_moving(current_order, updated_order_chan, internal_order_list, internal_order_list_chan, external_order_list, external_order_list_chan)
			elev_state = Door_open
		case Door_open:
			event_door_open(updated_order_chan, current_order, internal_order_list, external_order_list)
			fmt.Println("Out of Door_open.")
			elev_state = Idle
		case Stuck:
			event_stuck()
			elev_state = Idle
		}/*
		select{
		case elev_state := <- elev_state_chan:
			if elev_state == Idle {
				time.Sleep(100 * time.Millisecond)
				current_order, internal_order_list, external_order_list = event_idle(new_order_bool_chan, internal_order_list_chan, external_order_list_chan)
				go Elev_state_to_chan(Moving, elev_state_chan)
			}else if elev_state == Moving {
				event_moving(current_order, updated_order_chan, internal_order_list, internal_order_list_chan, external_order_list, external_order_list_chan)
				go Elev_state_to_chan(Door_open, elev_state_chan)
			}else if elev_state == Door_open {
				event_door_open(updated_order_chan, current_order)
				fmt.Println("Out of Door_open.")
				go Elev_state_to_chan(Idle, elev_state_chan)
			}else if elev_state == Stuck {
				event_stuck()
				go Elev_state_to_chan(Idle, elev_state_chan)
			}
		}*/
	}
}

func event_idle(new_order_bool_chan chan bool, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order) (queue.Order, [global.NUM_INTERNAL_ORDERS]queue.Order, [global.NUM_GLOBAL_ORDERS]queue.Order) {
	fmt.Println("Running event: Idle.")

	var current_order queue.Order
	//-- må endre til å ta inn nåværende liste og ikke starte med en tom en
	//var internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order
	//var external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order
	internal_order_list := <-internal_order_list_chan
	external_order_list := <-external_order_list_chan
	go queue.Internal_order_list_to_channel(internal_order_list, internal_order_list_chan)
	go queue.External_order_list_to_channel(external_order_list, external_order_list_chan)

	// Check if there is an order in one of the lists


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
			return current_order, internal_order_list, external_order_list
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
	elevator_to_floor(floor, order_list, updated_order_chan, current_order, internal_order_list, external_order_list)
}

func event_door_open(updated_order_chan chan queue.Order, current_order queue.Order, internal_order_list[global.NUM_INTERNAL_ORDERS] queue.Order, external_order_list[global.NUM_GLOBAL_ORDERS] queue.Order ) {
	fmt.Println("Running event: Door open.")

	// Open door
	driver.Open_door()
	fmt.Println("Door opened.")
	driver.Set_button_lamp(current_order.Button, current_order.Floor, global.OFF) //-- can be moved to before open door
	fmt.Println("Door open lamp set on.")
	var also_delete_order queue.Order

	// Set order state to finished

if current_order.Button != global.BUTTON_COMMAND{
	for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
		if internal_order_list[i].Floor == current_order.Floor {
			also_delete_order = internal_order_list[i]
		}
	}
}else{
	for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
		if external_order_list[i].Order_state != queue.Inactive {
			also_delete_order = external_order_list[i]
		}
	}
}
 current_order.Order_state = queue.Finished
 also_delete_order.Order_state = queue.Finished


	fmt.Println("Current order state set to finished, current_order: ", current_order)
	go queue.Order_to_updated_order_chan(current_order, updated_order_chan)
	go queue.Order_to_updated_order_chan(also_delete_order, updated_order_chan)
	fmt.Println("Order sent on updated order chan.")
	time.Sleep(6*time.Second)
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

func elevator_to_floor(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order, updated_order_chan chan queue.Order, current_order queue.Order, internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order, external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order) {
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
	current_floor_int := driver.Get_floor_sensor_signal()
	current_floor := driver.Floor_int_to_floor_t(current_floor_int)
	floor_int := driver.Floor_t_to_floor_int(floor)
	fmt.Println(current_floor_int, floor_int)

	if current_floor_int < floor_int {
		fmt.Println("Going up.")
		driver.Set_motor_direction(global.DIR_UP)

		for driver.Get_floor_sensor_signal() != floor_int {
			current_floor = driver.Floor_int_to_floor_t(driver.Get_floor_sensor_signal())

			// When arriving at any floor, check for order
			if driver.Get_floor_sensor_signal() != -1 {
				driver.Set_floor_indicator_lamp(floor)
				pick_up_order_on_the_way(current_floor, order_list, updated_order_chan, current_order, internal_order_list, external_order_list)
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
				driver.Set_floor_indicator_lamp(floor)
				pick_up_order_on_the_way(current_floor, order_list, updated_order_chan, current_order, internal_order_list, external_order_list)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}

	// Stop when at desired floor
	driver.Set_motor_direction(global.DIR_STOP)
}

func pick_up_order_on_the_way(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order, updated_order_chan chan queue.Order, current_order queue.Order, internal_order_list [global.NUM_INTERNAL_ORDERS]queue.Order, external_order_list [global.NUM_GLOBAL_ORDERS]queue.Order) {
	// If the elevator has an order in this floor, stop and take this order
	order_in_floor, on_the_way_order := check_if_order_in_floor(floor, order_list)
	if order_in_floor {
		driver.Set_motor_direction(global.DIR_STOP)
		time.Sleep(100*time.Millisecond) // it loops when running event_door_open, using the same order over and over again
		event_door_open(updated_order_chan, on_the_way_order, internal_order_list, external_order_list)
	}
}

func check_if_order_in_floor(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order) (bool, queue.Order) {
	for i := 0; i < global.NUM_ORDERS; i++ {
		if order_list[i].Floor == floor && order_list[i].Order_state != queue.Inactive {
			return true, order_list[i]
		}
	}
	return false, order_list[0]
}

/*
// From Morten Fyhn (github: mortenfyhn/TTJ4145-Heis/Lift/src/fsm/timer.go)
func door_timer(timeout chan<- bool, reset <-chan bool) {
	const ddoor_open_time = 3 * time.Second
	timer := time.NewTimer(0)
	timer.Stop()

	for {
		select {
		case <-reset:
			timer.Reset(door_open_time)
		case <-timer.C:
			timer.Stop()
			timeout <- true
		}
	}
}*/
