package queue

import(
  "global"
  "driver"
)

// order states
const (
	inactive = int = iota
	added
	assigned
	ready
	active
)

type order struct{
  	button global.button_t
	floor global.floor_t
	order_state int
	assigned_to assigned_t
}

type elev_info struct{
	elev_ip int
	elev_last_floor global.floor_t
	elev_dir global.motor_direction_t
	elev_state int
}

// declare variables
var order_state int

// initial values
func Init_queue(){
	order_state = inactive
}

// we want lists of length [n] and type order
global_order_list = [6] order
internal_order_list = [4] order
my_order_list = [10] order

// ------
	
//Global_order is the order queue the master is sending out, and is the order that should be acted out from
type Global_order struct{
		floor floor_t
		button button_t
		assigned_to int
    //order_state order_state
}

//When a slave receives a new order, they should send it to the master like this
type New_order struct{
			floor floor_t
			button button_t
			//elev_state elev_state
			elev_position int
			my_IP int
}

type my_orders struct {

}

var Global_order_matrix = []
