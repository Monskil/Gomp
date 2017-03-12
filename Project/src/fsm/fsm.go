package fsm

import (
  "global"
  "queue"
  "driver"
  "fmt"
  "time"
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


func State_handler(updated_order_chan chan queue.Order, global_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order ){
	fmt.Println("State handler is up and running")
	elev_state := Idle
	current_order_chan := make(chan queue.Order)
	var current_order queue.Order
	for {
		switch elev_state{
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

func event_idle(global_order_list_chan chan [global.NUM_GLOBAL_ORDERS]queue.Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]queue.Order, current_order_chan chan queue.Order)queue.Order{
	fmt.Println("Running event: Idle.")

	var current_order queue.Order
	// -- burde ta inn nåværende liste og ikke starte med en tom en
	var internal_order_list  [global.NUM_INTERNAL_ORDERS]queue.Order
	var global_order_list  [global.NUM_GLOBAL_ORDERS]queue.Order //hvis vi går inn i idle med full liste skjer det ingenting før ny ordre kommer
	for{
		for i := 0 ; i < global.NUM_INTERNAL_ORDERS ; i++ {
			if internal_order_list[i].Order_state != queue.Inactive {
				current_order = internal_order_list[i]
				return current_order
			}
		}
		for i := 0 ; i < global.NUM_GLOBAL_ORDERS ; i++ {
			if global_order_list[i].Order_state != queue.Inactive {
				current_order = global_order_list[i]
				return current_order
			}
		}

		select{
		case catch_internal_list :=<- internal_order_list_chan:
			internal_order_list = catch_internal_list
		case  catch_global_list :=<- global_order_list_chan:
			global_order_list = catch_global_list	
		}
	}
}


func event_moving(updated_order_chan chan queue.Order, current_order queue.Order){
	fmt.Println("Running event: Moving.")

	// Make order list to test, should be sent through channel
	var order_list [global.NUM_ORDERS]queue.Order
	order_list[2].Floor = global.FLOOR_2
	fmt.Println("My order list is: ", order_list)

	current_order.Order_state = queue.Executing
	updated_order_chan <- current_order
	//driver.Elevator_to_floor_direct(current_order.Floor)
	floor := current_order.Floor
	elevator_to_floor(floor, order_list, updated_order_chan, current_order)

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

func event_door_open(updated_order_chan chan queue.Order, current_order queue.Order){
	fmt.Println("Running event: Door open.")
	driver.Open_door()
	fmt.Println(driver.Get_floor_sensor_signal())
	current_order.Order_state = queue.Finished
	updated_order_chan <- current_order
}

func elevator_to_floor(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order,updated_order_chan chan queue.Order, current_order queue.Order){
	current_floor_int := driver.Get_floor_sensor_signal()
	current_floor := driver.Floor_int_to_floor_t(current_floor_int)

	floor_int := driver.Floor_t_to_floor_int(floor)

	// If the elevator is between two floors
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
	if current_floor_int < floor_int {
		fmt.Println("Going up.")
		for driver.Get_floor_sensor_signal() != floor_int {
			driver.Set_motor_direction(global.DIR_UP)
			current_floor = driver.Floor_int_to_floor_t(driver.Get_floor_sensor_signal())
			
			// When we arrive at any floor, check for order
			if driver.Get_floor_sensor_signal() != -1 {
				pick_up_order_on_the_way(current_floor, order_list, updated_order_chan, current_order)
				time.Sleep(10*time.Millisecond)
			}
		}
	} else if current_floor_int > floor_int {
		fmt.Println("Going down.")
		for driver.Get_floor_sensor_signal() != floor_int {
			driver.Set_motor_direction(global.DIR_DOWN)
			current_floor = driver.Floor_int_to_floor_t(driver.Get_floor_sensor_signal())
			
			// When we arrive at any floor, check for order
			if driver.Get_floor_sensor_signal() != -1 {
				pick_up_order_on_the_way(current_floor, order_list, updated_order_chan, current_order)
				time.Sleep(10*time.Millisecond)
			}
		}
	}
	// Stop when at desired floor
	driver.Set_motor_direction(global.DIR_STOP)
}

func pick_up_order_on_the_way(floor global.Floor_t, order_list [global.NUM_ORDERS]queue.Order, updated_order_chan chan queue.Order, current_order queue.Order){
	// If the elevator has an order in this floor, stop and take this order
	order_in_floor := queue.Check_if_order_in_floor(floor, order_list)
	if order_in_floor {
		driver.Set_motor_direction(global.DIR_STOP)
		event_door_open(updated_order_chan, current_order)
	}
}
