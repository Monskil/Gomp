package queue

import (
	"global"
)

//compares the cost for the n elevators
// should take in num elevators online

//-- must fix elevator [3]stuff
func compare_cost(elevator [global.NUM_ELEV]Elev_info, num_elevators_online int) Elev_info {

	lowest_cost := 100 // Calculate_cost(elevator[0])
	var best_suited_elevator Elev_info

	for i := 0; i < num_elevators_online; i++ {
		cost := Calculate_cost(elevator[i].Order_list, elevator[i].Elev_dir, elevator[i].Elev_destination_floor, elevator[i].Elev_last_floor)

		if cost == -2 {
			best_suited_elevator = elevator[i]
			break
		} else if cost < lowest_cost {
			best_suited_elevator = elevator[i]
			lowest_cost = cost
		}
	}
	return best_suited_elevator

}

//-- calculate cost need all of this information
//order_list global.NUM_ORDERS]Order, direction global.Motor_direction_t, destination_floor global.Floor_t, current_floor global.Floor_t
//have to get this from somewhere

//Takes in the IP for each elevator, so we can get the "order-list" for each elevator
//Then we can calculates the cost for each elevator and compare them
//Assigning the order to the elevator with the lowest cost
