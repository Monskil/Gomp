// The elevator has three states:
// Idle: not moving, waiting for orders
// Moving: moving, handling order
// Door open: at a floor with the door open, finishing order
//
// The orders have five states:
// Nonactive: no order
// Order added: the order has been added to the list, but not yet assigned
// Assigned: the order has been assigned
// Ready: verified that order is known at assigned elevator, waiting in line to be executed
// Active: the order is being executed
// Finished: the order is finished
//
// We have x events:
// New order: a new order is received
// Floor reached: desired floor is reached
// Door closed: the door goes from open to closed

package fsm

import (
	"fmt"
)

const (
	idle int = iota
	moving
	door_open
)

const (
	nonactive = int = iota
	order_added 
	assigned
	ready
	active
	finished
)

var elev_state int
//var order_state int
var floor int
var dir int

type Channels struct {
	New_order chan bool
	Floor_reached chan int
	Door_closed chan bool

	Motor_dir chan int
	Floor_lamp chan int
	Door_lamp chan int
}

func Init(channel Channels){
	elev_state = idle
	dir = DIR_STOP
	floor = FLOOR_1

	fmt.Println("FMS init done.")
}
