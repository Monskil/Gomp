// The elevator has 3 states:
// Idle: not moving, waiting for orders
// Moving: moving, handling order
// Door open: at a floor with the door open, finishing order
// Stuck: if you've been moving for more than 5 seconds (have timer on moving) you're stuck, when reaching floor -> not stuck anymore
//
// We have 3 events:
// New order: a new order is received
// Floor reached: desired floor is reached
// Door closed: the door goes from open to closed

package fsm

import (
	"fmt"
	"global"
	"ordermanager"
	"driver"
	"queue"
)

// Elevator states
const (
	Idle int = iota
	Moving
	Door_open
	Stuck
)

// Declare variables
var elev_state int
var floor global.Floor_t
var dir global.Motor_direction_t

// Make channels
type Channels struct {
	// Channels triggering events
	New_order chan bool
	Floor_reached chan int
	Door_close chan bool
	
	// Channels setting values
	Motor_dir chan global.Motor_direction_t
	Floor_lamp chan global.Floor_t
	Door_lamp chan int
}

// Initial values
func Init(){
	elev_state = Idle
	dir = global.DIR_STOP
	floor = global.FLOOR_1

	fmt.Println("FMS init done.")
}

// Wait for signals -> run events
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

// Event: new order
func event_new_order(channel Channels){
	fmt.Println("Running event: new order.")

	switch elev_state {
	case Idle:
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
	case Door_open:
		// if you are at the correct floor:
		//	order_state = finished
		
		// should be add some seconds to the timer in open door?
	default:
		// if not valid state
	}
}

// Event: floor reached
func event_floor_reached(channel Channels, floor global.Floor_t){
	fmt.Println("Running event: floor reached.")
	
	// Turn on floor lamp
	channel.Floor_lamp <- floor
	
	order := False
	switch elev_state {
	case Moving:
		// Check if the elevator has a order at this floor
		//order = ordermanager.Check_if_order_at_floor(floor global.Floor_t, elev global.Assigned_t)
		// if order:
		//	dir = global.DIR_STOP
		//	channel.Motor_dir <- dir
		// 	driver.Open_door()
		// 	elev_state = door_open
		//      -- Change state of the order belonging to this elevetor and floor (master-task?)       
		// 	--(order_state = finished)
		//	-- for all buttons, set to finished:
		//	queue.Update_order_state(button global.Button_t, floor global.Floor_t, state queue.Order_state, elev global.Assigned_t)
		// 	Update_order_state_to_finished_whole_floor_all_elev(floor)
		
		default:
		// if not valid state

}

// Event: door close
func event_door_close(channel Channels){
	fmt.Println("Running event: door close.")
	
	switch elev_state {
	case Door_open:
		// turn off door lamp
		channel.Door_lamp <- false
		
		// check for next order:
		
		//dir = direction of the next order
		
		// set motor direction
		// - hvor kommer dir fra?
		channel.Motor_dir <- dir
		
		// set elevator state
		if dir == global.DIR_STOP{
			elev_state = Idle
		} else {
			elev_state = Moving
			// order_state = active
		}
		
	default:
		// if not valid state
	}
}

// master function? om master fikser dette så blir det oppdatert for alle
func Update_order_state_to_finished_whole_floor_all_elev(floor global.Floor_t){
	// for all buttons
	// for all elev
	queue.Update_order_state(button, floor, finished, elev)
}
		
	
