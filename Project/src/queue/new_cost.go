package queue

import (
        "global"        
)

//pseudo code for new cost code

order_floor := order_from_panel //getting the order from panel. Is this a number?

//Checking if an elevator is free and just waiting for an order. 
//If thats the case, we should delegate the order to that elevator
func Check_if_free_elev(global_order_list, order) { 
	for i := 0; i = NUM_ORDERS; i++{
		if global_order_list[i] = 0 {
			global_order_list[4] = order //Element 4 er vel første eksterne element i lista
		}	
	}
}

//Simulates the elevators journey and calculates the cost
func Calculate_cost(global_order_list, internal_order_list, order){
	cost := 0

	for i := 0; i = NUM_ORDERS; i++ {
		current_floor = internal_order_list.floor 
		button = internal_order_list.button
		direction = determine_direction()

		//All cost functions
		cost += direction_cost(direction,)
		cost += stop_cost(direction,)
		cost += floor_cost(direction,)
	}
	
	return cost
}


func determine direction(curr_floor, destination_floor) Motor_direction_t {
	if(curr_floor > destination_floor){
		return DIR_DOWN
	} else {
		return DIR_UP
	}
}
	
				
//Calculates the cost based on the direction. Adds +3 for wrong dir and -1 for right dir
//Must take in current floor, destination floor and orderd floor
func direction_cost(direction) {
	direction_cost := 0
	
	switch direction{
		case DIR_DOWN:
		if (order_floor < curr_floor) {
			//Elevator is going down, destination is lower than current floor 
			direction cost -1
		} 
		else {
			//Elevator going down, destination is higher than current floor
			direction cost = 3
		}
		
		case DIR_UP:
		if (order_floor > curr_floor) {
			//Elevator going up, destination is higher than current floor
			direction cost -1
		} else {
			//Elevator going up, destination is lower than current floor
			direction cost = 3
		}
	}
	
	return direction_cost
}

//Calculates the cost based on the distance between the elevator and the destination floor.
//Adds +2 for each floor it passes
func floor_cost() {
	floor_cost := 0
	if (curr_floor < order_floor) {
		floor_cost = 2*(order_floor - curr_floor - 1)
	} else {
		floor_cost = (-2)*(order_floor - curr_floor + 1)
	}
	return floor_cost
}

//Calculates the cost based on stops. Adds +2 for each time it stops
//Should take in current floor, ordered floor and if a button is pressed
func stop_cost() {
	stop_cost := 0
	//If the elevator is going down, but stops on the way down
	if direction == DIR_DOWN {
		switch {
			case ((curr_floor - order_floor) == 3 && (button == BUTTON_COMMAND2 || button == BUTTON_DOWN3): 
				stop_cost += 2
				fallthrough	
			case ( (curr_floor - order_floor) == 3) && (button == BUTTON_COMMAND3):
				stop_cost += 2
				fallthrough
			case ( (curr_floor - order_floor) == 2) && (button == BUTTON_COMMAND2):
				stop_cost += 2
				fallthrough
			case ( (curr_floor - order_floor) == 2) && (button == BUTTON_COMMAND3):
				stop_cost += 2
		} else if direction == DIR_UP { //If elevator is going up, but stops on the way up
		switch {
			case ( (curr_floor - order_floor) == -3) && (button == BUTTON_COMMAND2):
				stop_cost += 2
				fallthrough
			case ( (curr_floor - order_floor) == -3) && (button == BUTTON_COMMAND3):
				stop_cost += 2
				fallthrough
			case ( (curr_floor - order_floor) == -2) && (button == BUTTON_COMMAND2):
				stop_cost += 2
				fallthrough
			case ( (curr_floor - order_floor) == -2) && (button == BUTTON_COMMAND3):
				stop_cost += 2
	}
		
	return stop_cost
}

func valid_floor_cost(){
	valid_floor_cost := 0
	if (curr_floor = -1){
		valid_floor_cost += 1
		return valid_floor_cost
	}
}
