// The elevator has 3 states:
// Idle: not moving, waiting for orders
// Moving: moving, handling order
// Door open: at a floor with the door open, finishing order
//
// We have 3 events:
// New order: a new order is received
// Floor reached: desired floor is reached
// Door closed: the door goes from open to closed

package fsm

import (
	"fmt"
	"global"
)

// elevator states
const (
	idle int = iota
	moving
	door_open
)

// ---- moved to queue ----
// order states
const (
	inactive = int = iota
	added
	assigned
	ready
	active
)

// order state management will be fixed later
//var order_state int
--------

// declare variables
var elev_state int
var floor global.floor_t
var dir global.motor_direction_t

// make channels
type Channels struct {
	// channels triggering events
	New_order chan bool
	Floor_reached chan int
	Door_close chan bool
	
	// channels setting values
	Motor_dir chan global.motor_direction_t
	Floor_lamp chan global.floor_t
	Door_lamp chan int
}

// initial values
func Init(){
	elev_state = idle
	dir = DIR_STOP
	floor = FLOOR_1

	fmt.Println("FMS init done.")
}

// wait for signals -> run events
func run(channel Channels){
	for{
		select{
		case <-channel.New_order:
			event_new_order(channel)
		case floor := <- channel.Floor_reached:
			event_floor_reached(channel, floor)
		case <- channel.Door_close:
			event_door_close(channel)
		}
	}
}

// event: new order
func event_new_order(channel Channels){
	fmt.Println("Event: new order.")

	switch elev_state {
	case idle:
		// get direction of the next order
		//dir = direction of the next order
		// if you are at the correct floor:
		//	open door
		//	elev_state = door_open
		//	order_state = finished
		// else:
		// 	channel.Motor_dir <- dir
		//	elev_state = moving
		//	order_state = active
	case door_open:
		// if you are at the correct floor:
		//	order_state = finished
		
		// should be add some seconds to the timer in open door?
	default:
		// if not valid state
	}
}

// event: floor reached
func event_floor_reached(channel Channels, floor def.floor_t){
	fmt.Println("Event: floor reached.")
	
	// turn on floor lamp
	channel.Floor_lamp <- floor
	
	switch elev_state {
	case moving:
		// check if order at this floor
		// if yes:
		//	dir = def.DIR_STOP
		//	channel.Motor_dir <- dir
		// 	open door
		// 	elev_state = door_open
		// 	order_state = finished
	default:
		// if not valid state

}

// event: door close
func event_door_close(channel Channels){
	fmt.Println("Event: door close.")
	
	switch elev_state {
	case door_open:
		// turn off door lamp
		channel.Door_lamp <- false
		
		// check for next order:
		
		//dir = direction of the next order
		
		// set motor direction
		channel.Motor_dir <- dir
		
		// set elevator state
		if dir == def.DIR_STOP{
			elev_state = idle
		} else {
			elev_state = moving
			// order_state = active
		}
		
	default:
		// if not valid state
}
