package queue


//compares the cost for the 3 elevators
func compare_cost() {

	if ((cost_1 < cost_2) && (cost_1 < cost_3)) {
		elevator 1 gets it
	} else if ((cost_2 < cost_1) && (cost_2 < cost_3)) {
		elevator 2 gets it
	} else if ((cost_3 < cost_1) && (cost_3 < cost_2)) {
		elevator 3 gets it
	} else {
		elevator with lowest IP gets it
	}

}

//Takes in the IP for each elevator, so we can get the "order-list" for each elevator
//Then we can calculates the cost for each elevator and compare them
//Assigning the order to the elevator with the lowest cost

//calculates the cost for elevator 1
func cost_elevator_1(order_list [global.NUM_ORDERS].queue.Order, direction global.Motion_direction_t, destination_floor global.Floor_t, current_floor global.Floor_t) int {

	cost_1 := Calculate_cost(order_list, direction, destination_floor, current_floor)

}

//calculates the cost for elevator 2
func cost_elevator_2(order_list [global.NUM_ORDERS].queue.Order, direction global.Motion_direction_t, destination_floor global.Floor_t, current_floor global.Floor_t) int {

	cost_2 := Calculate_cost(order_list, direction, destination_floor, current_floor)

}

//calculates the cost for elevator 3
func cost_elevator_3(order_list [global.NUM_ORDERS].queue.Order, direction global.Motion_direction_t, destination_floor global.Floor_t, current_floor global.Floor_t) int {

	cost_3 := Calculate_cost(order_list, direction, destination_floor, current_floor)

}
