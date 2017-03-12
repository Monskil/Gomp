package queue

import (
        "global"        
)

//pseudo code for new cost code

order_floor := order_from_panel //getting the order from panel. Is this a number?

//Checking if an elevator is free and just waiting for an order. 
//If thats the case, we should delegate the order to that elevator
func elevator_is_idle(order_list [global.NUM_ORDERS]queue.Order) bool{ 
	for i := 0; i < NUM_ORDERS; i++{
		if order_list[i].Order_state != queue.Inactive {
			return false
		}
	}
	return true
}

func orders_in_list_cost(order_list [global.NUM_ORDERS]queue.Order)int{ 
	cost := 0
	for i := 0; i < NUM_ORDERS; i++{
		if order_list[i].Order_state != queue.Inactive {
			cost += 2
		}
	}
	return cost
}

//If elevator is idle returns true do this
global_order_list[4] = order //Element 4 er vel første eksterne element i lista


//Simulates the elevators journey and calculates the cost
//func Calculate_cost(global_order_list, internal_order_list, order){
//	cost := 0
//	previous_floor = Forrige gyldig etasje
//	current_floor = Etasjen/posisjonen heisen er i
//	button = Hva som kommer inn fra knappen
//	for i := 0; i = NUM_ORDERS; i++ {		
//		
//		direction = determine_direction()

//		cost += direction_cost(direction,)
//		cost += stop_cost(direction,)
//		cost += floor_cost(direction,)
//		cost += between_floor_cost()
//	}
	
//	return cost
//}


func determine_direction(previous_floor global.Floor_t, destination_floor global.Floor_t) global.Motor_direction_t {
	if(previous_floor > destination_floor){
		return DIR_DOWN
	} else {
		return DIR_UP
	}
}
	
				
//Calculates the cost based on the direction. Adds +3 for wrong dir and -1 for right dir
//Must take in current floor, destination floor and orderd floor
func direction_cost(direction global.Motion_direction_t, destination_floor global.Floor_t, current_floor global.Floot_t) int {
	direction_cost := 0
	
	switch direction{
		case DIR_DOWN:
		if (destination_floor < current_floor) {
			//Elevator is going down, destination is lower than current floor 
			direction cost -1
		} 
		else {
			//Elevator going down, destination is higher than current floor
			direction cost = 3
		}
		
		case DIR_UP:
		if (order_floor > current_floor) {
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
func floor_cost(current_floor global.Floor_t, order_floor global.Floor_t) {
	floor_cost := 0
	
	if (current_floor < order_floor) {
		floor_cost = 2*(order_floor - current_floor - 1)
	} else {
		floor_cost = (-2)*(order_floor - current_floor + 1)
	}
	
	return floor_cost
}

//Calculates the cost based on stops. Adds +2 for each time it stops
//Should take in current floor, ordered floor and if a button is pressed
func stop_cost(direction global.Motor_direction_t, previous_floor global.Floor_t, order_floor globabl.Floor_t) {
	stop_cost := 0
	//If the elevator is going down, but stops on the way down
	if direction == DIR_DOWN {
		switch {
			case ( (previous_floor - order_floor) == 3 && (button == BUTTON_COMMAND3 || button == BUTTON_DOWN3): 
				stop_cost += 2
				fallthrough	
			case ( (previous_floor - order_floor) == 3) && (button == BUTTON_COMMAND2 || button == BUTTON_DOWN2):
				stop_cost += 2
			      	fallthrough
			case ( (previous_floor - order_floor) == 2) && (button == BUTTON_COMMAND2 || button == BUTTON_DOWN2):
				stop_cost += 2
				fallthrough
			case ( (previous_floor - order_floor) == 2) && (button == BUTTON_COMMAND3 || button == BUTTON_DOWN3):
				stop_cost += 2
			      
		} else if direction == DIR_UP { //If elevator is going up, but stops on the way up
		switch {
			case ( (previous_floor - order_floor) == -3) && (button == BUTTON_COMMAND2 || button == BUTTON_UP2):
				stop_cost += 2
				fallthrough
			case ( (previous_floor - order_floor) == -3) && (button == BUTTON_COMMAND3 || button == BUTTON_UP3):
				stop_cost += 2
				fallthrough
			case ( (previous_floor - order_floor) == -2) && (button == BUTTON_COMMAND2 || button == BUTTON_UP2):
				stop_cost += 2
				fallthrough
			case ( (previous_floor - order_floor) == -2) && (button == BUTTON_COMMAND3 || button == BUTTON_UP3):
				stop_cost += 2
	}
		
	return stop_cost
}

//Calculates cost if the elevator starts in an invalid position. Adding 1 if so
func between_floor_cost(){
	between_floor_cost := 0
	
	if (curr_floor == -1){
		between_floor_cost = 1
		return between_floor_cost
	}
}
