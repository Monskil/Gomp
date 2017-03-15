package queue

import (
	"driver"
	"global"
)

func Compare_cost(elevators [global.NUM_ELEV]Elev_info, num_elevators_online int) Elev_info {
	lowest_cost := 100
	var best_suited_elevator Elev_info

	for i := 0; i < num_elevators_online; i++ {
		cost := Calculate_cost(elevators[i].Order_list, elevators[i].Elev_dir, elevators[i].Elev_destination_floor, elevators[i].Elev_last_floor)

		if cost == -2 {
			best_suited_elevator = elevators[i]
			break
		} else if cost < lowest_cost {
			best_suited_elevator = elevators[i]
			lowest_cost = cost
		}
	}
	return best_suited_elevator
}

func calculate_cost(order_list [global.NUM_ORDERS]Order, direction global.Motor_direction_t, destination_floor global.Floor_t, current_floor global.Floor_t) int {
	cost := 0

	if elevator_is_idle(order_list) {
		cost = -2
		return cost
	} else {
		cost += order_cost(order_list)
		cost += direction_cost(direction, destination_floor, current_floor)
		cost += floor_cost(current_floor, destination_floor)
	}
	return cost
}

// Checking if elevator is free and just waiting for an order
func elevator_is_idle(order_list [global.NUM_ORDERS]Order) bool {
	for i := 0; i < global.NUM_ORDERS; i++ {
		if order_list[i].Order_state != Inactive {
			return false
		}
	}
	return true
}

// Add 2 points for each order in list
func order_cost(order_list [global.NUM_ORDERS]Order) int {
	order_cost := 0
	for i := 0; i < global.NUM_ORDERS; i++ {
		if order_list[i].Order_state != Inactive {
			order_cost += 2
		}
	}
	return order_cost
}

// Add 3 points for wrong direction and -1 for right direction
func direction_cost(direction global.Motor_direction_t, destination_floor global.Floor_t, current_floor global.Floor_t) int {
	direction_cost := 0

	switch direction {
	case global.DIR_DOWN:
		if destination_floor < current_floor {
			//Elevator is going down, destination is lower than current floor
			direction_cost = -1
		} else {
			//Elevator going down, destination is higher than current floor
			direction_cost = 3
		}

	case global.DIR_UP:
		if destination_floor > current_floor {
			//Elevator going up, destination is higher than current floor
			direction_cost = -1
		} else {
			//Elevator going up, destination is lower than current floor
			direction_cost = 3
		}
	}

	return direction_cost
}

// Add 2 points for each floor between the elevator and the order
func floor_cost(current_floor global.Floor_t, order_floor global.Floor_t) int {
	floor_cost := 0

	if current_floor < order_floor {
		floor_cost = 2 * (driver.Floor_t_to_floor_int(order_floor) - driver.Floor_t_to_floor_int(current_floor) - 1)
	} else {
		floor_cost = (-2) * (driver.Floor_t_to_floor_int(order_floor) - driver.Floor_t_to_floor_int(current_floor) + 1)
	}

	return floor_cost
}
