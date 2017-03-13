package queue


//compares the cost for the n elevators
// should take in num elevators online
func compare_cost(elevator[num_elevators_online] Elev_info, num_elevators_online int) Elev_info.Elev_ip {

	lowest_cost := 100 // Calculate_cost(elevator[0])
	var best_suited_elevator Elev_info


	for i:= 0 ; i < num_elevators_online ; i++ {
		cost := Calculate_cost(elevator[i])
		
		if cost == -2 {
			best_suited_elevator = elevator[i]
			break
		} else if cost < lowest_cost { 
			best_suited_elevator = elevator[i]
			lowest_cost = cost
		}
	}
	return best_suited_elevator.Elev_ip

}

//Takes in the IP for each elevator, so we can get the "order-list" for each elevator
//Then we can calculates the cost for each elevator and compare them
//Assigning the order to the elevator with the lowest cost
