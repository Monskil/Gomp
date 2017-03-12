package queue

import (
        "global"
        
)

//pseudo code for new cost code


order_floor := order_from_panel //getting the order from panel. Is this a number?
cost := 0


//Checking if an elevator is free and just waiting for an order. 
//If thats the case, we should delegate the order to that elevator

func Check_if_free_elev(global_order_list, order) { 
	for i := 0; i = NUM_ORDERS; i++{
		if global_order_list[i] = 0 {

			global_order_list[4] = order //element 4 er vel første eksterne element i lista
		}	
	}
}

//Simulates the elevators journey and calculates the cost
func Calculate_cost(global_order_list, internal_order_list, order){
	for i := 0; i= NUM_ORDERS; i++ {
			curr_floor = internal_order_list.floor 
			button = internal_order_list.button
			//All cost functions
			direction_cost()
			stop_cost()
			floor_cost()

		}
	}
}

//Calculates the cost based on the direction. Adds +3 for wrong dir and -1 for right dir
//Must take in current floor, destination floor and orderd floor
func direction_cost() {
	direction_cost := 0
	switch {
		//Elevator is going down, destination is lower than current floor
		case (curr_floor > destination_floor) && (order_floor < curr_floor): 
			direction_cost = direction_cost - 1
		//Elevator going down, destination is higher than current floor
		case (curr_floor > destination_floor) && (order_floor > curr_floor):
			direction_cost = direction_cost + 3
		//Elevator going up, destination is higher than current floor
		case (curr_floor < destination_floor) && (order_floor > curr_floor):
			direction_cost = direction_cost - 1
		//Elevator going up, destination is lower than current floor
		case (curr_floor < destination_floor) && (order_floor < curr_floor):
			direction_cost = direction_cost + 3
	}
	return direction_cost

}

//Calculates the cost based on the distance between the elevator and the destination floor.
//Adds +2 for each floor it passes
func floor_cost() {
	if (curr_floor < order_floor) {
		cost = cost + 2*(order_floor - curr_floor - 1)
	} else {
		cost = cost + (-2)*(order_floor - curr_floor + 1)
	}
	return cost
}
//Calculates the cost based on stops. Adds +2 for each time it stops
func stop_cost() {
	//If the elevator is going down, but stops on way down
	if (curr_floor - order_floor > 1) {
		switch
		
		case (curr_floor - order_floor = 3) && (button = BUTTON_COMMAND2): //Is this the same as the internal button is pressed 2?
			cost = cost + 2
		case (curr_floor - order_floor = 3) && (button = BUTTON_COMMAND3):
			cost = cost + 2
		case (curr_floor - order_floor = 2) && (button = BUTTON_COMMAND2):
			cost = cost + 2
		case (curr_floor - order_floor = 2) && (button = BUTTON_COMMAND3):
			cost = cost + 2
		case 
			return cost
	}

	//If elevator is going up, but stops on way up
	if (curr_floor - order < -1) {
		switch
		case (curr_floor - order_floor = -3) && (button = BUTTON_COMMAND2):
			cost = cost + 2
		case (curr_floor - order_floor = -3) && (button = BUTTON_COMMAND3):
			cost = cost + 2
		case (curr_floor - order_floor = -2) && (button = BUTTON_COMMAND2):
			cost = cost + 2
		case (curr_floor - order_floor = -2) && (button = BUTTON_COMMAND3):
			cost = cost + 2
		case
			return cost
	}
	
}

func valid_floor_cost(){
	if (curr_floor = -1){
	cost = cost + 1
	return cost
	}
}
