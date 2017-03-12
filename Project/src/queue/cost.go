package queue

import (
        "global"      
        "queue"
)

// Checking if elevator is free and just waiting for an order 
func elevator_is_idle(order_list [global.NUM_ORDERS]queue.Order) bool{ 
	for i := 0; i < NUM_ORDERS; i++{
		if order_list[i].Order_state != queue.Inactive {
			return false
		}
	}
	return true
}

// Add 2 points for each order in list
func order_cost(order_list [global.NUM_ORDERS]queue.Order)int{ 
	order_cost := 0
	for i := 0; i < NUM_ORDERS; i++{
		if order_list[i].Order_state != queue.Inactive {
			order_cost += 2
		}
	}
	return order_cost
}
	
// Add 3 points for wrong direction and -1 for right direction		
func direction_cost(direction global.Motion_direction_t, destination_floor global.Floor_t, current_floor global.Floot_t) int {
	direction_cost := 0
	
	switch direction{
		case global.DIR_DOWN:
		if (destination_floor < current_floor) {
			//Elevator is going down, destination is lower than current floor 
			direction cost -1
		} 
		else {
			//Elevator going down, destination is higher than current floor
			direction cost = 3
		}
		
		case global.DIR_UP:
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

// Add 2 points for each floor between the elevator and the order
func floor_cost(current_floor global.Floor_t, order_floor global.Floor_t) {
	floor_cost := 0
	
	if (current_floor < order_floor) {
		floor_cost = 2*(order_floor - current_floor - 1)
	} else {
		floor_cost = (-2)*(order_floor - current_floor + 1)
	}
	
	return floor_cost
}

func Calculate_cost(order_list [global.NUM_ORDERS]queue.Order, direction global.Motion_direction_t, destination_floor global.Floor_t, current_floor global.Floot_t)int{
  cost := 0
  
  if elevator_is_idle(order_list [global.NUM_ORDERS]queue.Order){
    cost = -2
    return cost
  } else {
    cost += order_cost(order_list)
    cost += direction_cost(direction, destination_floor, current_floor)
    cost += floor_cost(current_floor, order_floor)
  }
  return cost
}
