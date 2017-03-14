package queue

import (
	"driver"
	"fmt"
	"global"
)

//Order states
const (
	Inactive  = iota //Noone has ordered this
	Active           //The order has been detected by button owner
	Assigned         //Master has assigned the order to someone, but it's not confirmed that someone noticed that they got it
	Ready            //Verified that order is known at assigned elevator, waiting in line to be executed
	Executing        //This order is now being executed, i.e is the first in line
	Finished         //Slave says the order is finished, master can delete it

)

type Order struct {
	Button      global.Button_t
	Floor       global.Floor_t
	Order_state int
	Assigned_to global.Assigned_t
}

type Elev_info struct {
	Elev_ip         string
	Elev_last_floor global.Floor_t
	Elev_dir        global.Motor_direction_t
	Elev_state      int
	//-- Can/should we put it here
	Order_list             [global.NUM_ORDERS]Order
	Elev_destination_floor global.Floor_t
}

// ---- Go functions to make channels love us again -----
func Internal_order_list_to_channel(internal_order_list [global.NUM_INTERNAL_ORDERS]Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	internal_order_list_chan <- internal_order_list
}

func External_order_list_to_channel(external_order_list [global.NUM_GLOBAL_ORDERS]Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order) {
	external_order_list_chan <- external_order_list
}

func Bool_to_new_order_channel(value bool, new_order_bool_chan chan bool) {
	new_order_bool_chan <- value
}

func Order_to_updated_order_chan(order Order, updated_order_chan chan Order) {
	updated_order_chan <- order
}

//----------------------------------------------------------

func Order_handler(new_order_bool_chan chan bool, new_order_chan chan Order, update_order_chan chan Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	fmt.Print("Running: Order handler. ")
	var external_order_list [global.NUM_GLOBAL_ORDERS]Order
	var internal_order_list [global.NUM_INTERNAL_ORDERS]Order

	// Put the empty lists on the channel so the channel is not empty
	go Internal_order_list_to_channel(internal_order_list, internal_order_list_chan)
	go External_order_list_to_channel(external_order_list, external_order_list_chan)

	for {
		select {
		case catch_new_order := <-new_order_chan:
			fmt.Println("Case: new button pressed")

			new_order := catch_new_order

			// Get order lists
			internal_order_list := <-internal_order_list_chan
			external_order_list := <-external_order_list_chan

			if new_order.Button == global.BUTTON_COMMAND {
				internal_order_list = Add_new_internal_order(new_order, internal_order_list)

				fmt.Println("Internal order added and sent.")
			} else {
				external_order_list = Add_new_external_order(new_order, external_order_list)

				fmt.Println("External order added and sent.")
			}
			// Push lists back to channels
			go Internal_order_list_to_channel(internal_order_list, internal_order_list_chan)
			go External_order_list_to_channel(external_order_list, external_order_list_chan)

			// Set button lamp
			driver.Set_button_lamp(new_order.Button, new_order.Floor, global.ON)

			// Let the world know we have a new order
			go Bool_to_new_order_channel(true, new_order_bool_chan)
			fmt.Println("New order bool chan <- true")

		case catch_update_order := <-update_order_chan:
			fmt.Println("Case: update order")

			// Get order lists
			internal_order_list := <-internal_order_list_chan
			external_order_list := <-external_order_list_chan

			update_order := catch_update_order

			// Update state
			internal_order_list, external_order_list = Update_state(update_order, internal_order_list, external_order_list)

			// Delete order if the order state is finished
			if update_order.Order_state == Finished {
				fmt.Println("Order is marked finished, update_order: ", update_order)
				internal_order_list, external_order_list = Delete_order(update_order, internal_order_list, external_order_list)
			}
			fmt.Println("Finished updating state.")

			// Push lists back to channels
			go Internal_order_list_to_channel(internal_order_list, internal_order_list_chan)
			go External_order_list_to_channel(external_order_list, external_order_list_chan)
			fmt.Println("Lists are now:", internal_order_list, external_order_list)

		}
	}
}
func Update_state(update_order Order, internal_order_list [global.NUM_INTERNAL_ORDERS]Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order) ([global.NUM_INTERNAL_ORDERS]Order, [global.NUM_GLOBAL_ORDERS]Order) {
	fmt.Println("Running Update_state")
	fmt.Println("My update_order: ", update_order)

	for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
		fmt.Println("Checking element ", i, " in internal list: ", internal_order_list[i])
		if update_order.Button == internal_order_list[i].Button && update_order.Floor == internal_order_list[i].Floor && internal_order_list[i].Order_state != Inactive {
			fmt.Println("Update internal order loop.")
			internal_order_list[i].Order_state = update_order.Order_state
			return internal_order_list, external_order_list
		}
	}
	for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
		fmt.Println("Checking element ", i, " in external list: ", external_order_list[i])
		if update_order.Button == external_order_list[i].Button && update_order.Floor == external_order_list[i].Floor && external_order_list[i].Order_state != Inactive {
			fmt.Println("Update external order loop.")
			external_order_list[i].Order_state = update_order.Order_state
			return internal_order_list, external_order_list
		}
	}

	fmt.Println("Error: State was not updated, update order: ", update_order)
	return internal_order_list, external_order_list
}

func Add_new_internal_order(new_order Order, internal_order_list [global.NUM_INTERNAL_ORDERS]Order) [global.NUM_INTERNAL_ORDERS]Order {
	new_order_floor := new_order.Floor

	for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
		if internal_order_list[i].Order_state == Inactive {
			internal_order_list[i] = new_order
			fmt.Println("New internal order was added!")
			return internal_order_list
		}
		if internal_order_list[i].Floor == new_order_floor {
			fmt.Println("The order is already in the internal order list.", internal_order_list[i])
			return internal_order_list
		}
	}
	fmt.Println("Error: No internal order was added.")
	return internal_order_list
}

func Add_new_external_order(new_order Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order) [global.NUM_GLOBAL_ORDERS]Order {
	new_order_floor := new_order.Floor
	new_order_button := new_order.Button

	for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
		if external_order_list[i].Order_state == Inactive {
			external_order_list[i] = new_order
			fmt.Println("New external order was added!")
			return external_order_list
		}
		if external_order_list[i].Floor == new_order_floor && external_order_list[i].Button == new_order_button {
			fmt.Println("The order is already in the global order list.", external_order_list[i])
			return external_order_list
		}
	}
	fmt.Println("Error: No external order was added.")
	return external_order_list
}

func Delete_order(updated_order Order, internal_order_list [global.NUM_INTERNAL_ORDERS]Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order) ([global.NUM_INTERNAL_ORDERS]Order, [global.NUM_GLOBAL_ORDERS]Order) {
	fmt.Println("Inside delete_order function.")
	if updated_order.Button == global.BUTTON_UP || updated_order.Button == global.BUTTON_DOWN {
		fmt.Println("Going to delete an external order.")
		external_order_list = Delete_external_order(updated_order, external_order_list)
	} else {
		fmt.Println("Going to delete an internal order.")
		internal_order_list = Delete_internal_order(updated_order, internal_order_list)
	}
	return internal_order_list, external_order_list
}

func Delete_internal_order(updated_order Order, internal_order_list [global.NUM_INTERNAL_ORDERS]Order) [global.NUM_INTERNAL_ORDERS]Order {
	// Delete all finished orders in the internal list by moving all the "later" orders one step forward,
	// make the last element in the list an "empty" order
	clean_order := Make_new_order(global.BUTTON_UP, global.FLOOR_1, Inactive, global.NONE)

	for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
		if internal_order_list[i].Order_state == Finished {
			fmt.Println("An internal order is marked finished.")
			for j := i; j < global.NUM_INTERNAL_ORDERS; j++ {
				if j < global.NUM_INTERNAL_ORDERS-1 {
					internal_order_list[j] = internal_order_list[j+1]
				} else if j == global.NUM_INTERNAL_ORDERS-1 {
					internal_order_list[j] = clean_order
				}
			}
		}
	}
	fmt.Println("Internal order deleted, internal order list is now: ", internal_order_list)
	return internal_order_list
}

func Delete_external_order(updated_order Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order) [global.NUM_GLOBAL_ORDERS]Order {
	// Delete all finished orders in the external list by moving all the "later" orders one step forward,
	// make the last element in the list an "empty" order
	clean_order := Make_new_order(global.BUTTON_UP, global.FLOOR_1, Inactive, global.NONE)

	for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
		if external_order_list[i].Order_state == Finished {
			fmt.Println("An external order is marked finished.")
			for j := i; j < global.NUM_ORDERS; j++ {
				if j < global.NUM_GLOBAL_ORDERS-1 {
					external_order_list[j] = external_order_list[j+1]
				} else if j == global.NUM_GLOBAL_ORDERS-1 {
					external_order_list[j] = clean_order
				}
			}
		}
	}
	fmt.Println("External order deleted, external order list is now: ", external_order_list)
	return external_order_list
}

func Make_my_order_list(internal_order_list [global.NUM_INTERNAL_ORDERS]Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order) [global.NUM_ORDERS]Order {
	// The first four elements are from the internal list,
	// the six last elements are from the external list

	var my_order_list [global.NUM_ORDERS]Order
	for i := 0; i < global.NUM_ORDERS; i++ {
		if i < global.NUM_INTERNAL_ORDERS {
			my_order_list[i] = internal_order_list[i]
		} else {
			my_order_list[i] = external_order_list[i-global.NUM_INTERNAL_ORDERS]
		}
	}
	return my_order_list
}

func Make_new_order(button global.Button_t, floor global.Floor_t, order_state int, assigned_to global.Assigned_t) Order {
	var new_order Order

	new_order.Button = button
	new_order.Floor = floor
	new_order.Order_state = order_state
	new_order.Assigned_to = assigned_to

	return new_order
}
