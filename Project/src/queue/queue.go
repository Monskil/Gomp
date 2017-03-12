package queue

import (
	"global"
	//"driver"
	"fmt"
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
}

var external_order_list [global.NUM_GLOBAL_ORDERS]Order
var internal_order_list [global.NUM_INTERNAL_ORDERS]Order

func Handle_orders(new_order_bool_chan chan bool, new_order_chan chan Order, updated_order_chan chan Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	fmt.Print("Running: Handle orders. ")
	for {
		select {
		case catch_new_order := <-new_order_chan:
			fmt.Println("Case: new button pressed")

			add_order(catch_new_order, external_order_list_chan, internal_order_list_chan)
			fmt.Println("Order is added in list")

			new_order_bool_chan <- true
			fmt.Println("New order bool chan <- true")

		case catch_updated_order := <-updated_order_chan:
			fmt.Println("Case: update order")
			updated_order := catch_updated_order
			Update_state(updated_order, external_order_list_chan, internal_order_list_chan)
			if updated_order.Order_state == Finished {
				fmt.Println("Order is finished")
				Delete_order(updated_order, internal_order_list_chan, external_order_list_chan)
			}
			//if updated_order.Order_state == Finished{
			//Delete_internal_order(updated_order, internal_order_list_chan)
			//}

		}
	}
}
func Update_state(updated_order Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	fmt.Println("Updates state before")
	fmt.Println("My update_order: ", updated_order)
	if updated_order.Button == global.BUTTON_UP || updated_order.Button == global.BUTTON_DOWN {
		fmt.Println("guk")
		//var external_order_list [global.NUM_GLOBAL_ORDERS]Order
		//external_order_list = <-external_order_list_chan
		fmt.Println("jkdil")
		external_order_list[1].Order_state = updated_order.Order_state
		external_order_list_chan <- external_order_list
	} else {
		internal_order_list := <-internal_order_list_chan
		internal_order_list[1].Order_state = updated_order.Order_state
		//internal_order_list_chan <- internal_order_list

	}
	fmt.Println("Updated the state")
}

func Add_new_internal_order(new_order Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	new_order_floor := new_order.Floor

	for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
		if internal_order_list[i].Order_state == Inactive {
			internal_order_list[i] = new_order
			fmt.Println("New internal order was added!")
			internal_order_list_chan <- internal_order_list
			fmt.Println("internal order list channel pushed to")
			break
		}
		if internal_order_list[i].Floor == new_order_floor {
			fmt.Println("The order is already in the internal order list.", internal_order_list[i])
			break

		}
	}
}

func Add_new_external_order(newButton Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order) {
	new_order_floor := newButton.Floor
	new_order_button := newButton.Button
	for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
		if external_order_list[i].Order_state == Inactive {
			external_order_list[i] = newButton
			fmt.Println("New external order was added!")
			external_order_list_chan <- external_order_list
			fmt.Println("Pushed to external order chan")
			break
		}
		if external_order_list[i].Floor == new_order_floor && external_order_list[i].Button == new_order_button {
			fmt.Println("The order is already in the global order list.", external_order_list[i])
			break
		}
	}

}

func add_order(new_order Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	fmt.Println("Inside add_order funccd..")

	if new_order.Button == global.BUTTON_UP || new_order.Button == global.BUTTON_DOWN {
		fmt.Println("Gikk inn i if'en")
		Add_new_external_order(new_order, external_order_list_chan)
	} else {
		Add_new_internal_order(new_order, internal_order_list_chan)
	}

}

func Delete_order(updated_order Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order, global_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order) {
	fmt.Println("in delete_order func")
	if updated_order.Button == global.BUTTON_UP || updated_order.Button == global.BUTTON_DOWN {
		fmt.Println("Det er en ekstern ordre som skal slettes")
		Delete_external_order(updated_order, global_order_list_chan)
	} else {
		fmt.Println("Det er en intern ordre som skal slettes")
		Delete_internal_order(updated_order, internal_order_list_chan)
	}
}

func Delete_internal_order(updated_order Order, internal_order_list_chan chan [global.NUM_INTERNAL_ORDERS]Order) {
	// delete all finished orders in the internal list by moving all the "later" orders one step forward
	// make the last element in the list an "empty" order
	//clean_order := Make_new_order(global.BUTTON_UP, global.FLOOR_1, Inactive, global.NONE)
	var clean_order Order
	clean_order.Button = updated_order.Button
	clean_order.Floor = updated_order.Floor
	clean_order.Order_state = Inactive
	clean_order.Assigned_to = global.NONE
	internal_order_list := <-internal_order_list_chan

	for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {
		fmt.Println("in the loop", internal_order_list[i].Order_state, Finished)
		if internal_order_list[i].Order_state == Finished {
			fmt.Println("An internal order is marked finished.")
			for j := i; j < global.NUM_INTERNAL_ORDERS; j++ {
				if j < global.NUM_INTERNAL_ORDERS-1 {
					fmt.Println("Moving order.")
					internal_order_list[j] = internal_order_list[j+1]
				} else if j == global.NUM_INTERNAL_ORDERS-1 {
					fmt.Println("Adding last clean order.")
					internal_order_list[j] = clean_order
				}
			}
		}
	}
	fmt.Println("almost finished")
	internal_order_list_chan <- internal_order_list
}

func Delete_external_order(updated_order Order, external_order_list_chan chan [global.NUM_GLOBAL_ORDERS]Order) {
	// delete all finished orders in the global list by moving all the "later" orders one step forward
	// make the last element in the list an "empty" order
	//clean_order := Make_new_order(global.BUTTON_UP, global.FLOOR_1, Inactive, global.NONE)
	var clean_order Order
	clean_order.Button = updated_order.Button
	clean_order.Floor = updated_order.Floor
	clean_order.Order_state = Inactive
	clean_order.Assigned_to = global.NONE

	for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
		if external_order_list[i].Order_state == Finished {
			fmt.Println("A global order is marked finished.")
			for j := i; j < global.NUM_ORDERS; j++ {
				if j < global.NUM_GLOBAL_ORDERS-1 {
					fmt.Println("Moving order.")
					external_order_list[j] = external_order_list[j+1]
				} else if j == global.NUM_GLOBAL_ORDERS-1 {
					fmt.Println("Adding last clean order.")
					external_order_list[j] = clean_order
				}
			}
		}
	}
	external_order_list_chan <- external_order_list
}

/*
func Make_my_order_list(internal_order_list [global.NUM_GLOBAL_ORDERS]Order, global_order_list [global.NUM_GLOBAL_ORDERS]Order) [global.NUM_ORDERS]Order {
  // add first the elements from the internal list
  // then the elements of the global list

  var my_order_list [global.NUM_ORDERS]Order
  for i := 0; i < global.NUM_ORDERS; i++ {
    if i < global.NUM_INTERNAL_ORDERS {
      my_order_list[i] = internal_order_list[i]
    } else {
      my_order_list[i] = global_order_list[i-global.NUM_INTERNAL_ORDERS]
    }
  }
  return my_order_list
}

func Make_new_order(button global.Button_t, floor global.Floor_t, order_state int, assigned_to global.Assigned_t) Order {
  // make new order with the elements you choose
  var new_order Order

  new_order.Button = button
  new_order.Floor = floor
  new_order.Order_state = order_state
  new_order.Assigned_to = assigned_to

  return new_order
}


*/

/*
 */

/*
func Add_new_global_order(new_order Order, global_order_list [global.NUM_GLOBAL_ORDERS]Order) [global.NUM_GLOBAL_ORDERS]Order {
  // add new order in the global list, if it's not already there
  new_order_floor := new_order.Floor
  new_order_button := new_order.Button

  for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {
    if global_order_list[i].Order_state == Inactive {
      global_order_list[i] = new_order
      fmt.Println("New order was added!")
      return global_order_list
    }
    if global_order_list[i].Floor == new_order_floor && global_order_list[i].Button == new_order_button {
      fmt.Println("The order is already in the global order list.")
      return global_order_list
    }
  }
  fmt.Println("Something went wrong, we didn't add any new orders.")
  return global_order_list
}*/
/*

// initial values
// and also some testing
func Init_queue() {
  fmt.Println("Helloo, I'm initializing the queue, yayy!")

  // make lists
  //var internal_order_list [global.NUM_INTERNAL_ORDERS]Order
  //var global_order_list [global.NUM_GLOBAL_ORDERS]Order
  //var my_order_list [global.NUM_ORDERS]Order

  // test example making orders/*
  /*
     var order1 = Make_new_order(global.BUTTON_UP, global.FLOOR_2, Finished, global.ELEV_2)
     var order2 = Make_new_order(global.BUTTON_COMMAND, global.FLOOR_1, Active, global.ELEV_3)
     var order3 = Make_new_order(global.BUTTON_UP, global.FLOOR_1, Active, global.ELEV_1)
     var order4 = Make_new_order(global.BUTTON_COMMAND, global.FLOOR_2, Active, global.ELEV_2)

     internal_order_list = Add_new_internal_order(order1, internal_order_list)
     internal_order_list = Add_new_internal_order(order2, internal_order_list)
     internal_order_list = Add_new_internal_order(order2, internal_order_list)

     global_order_list = Add_new_global_order(order2, global_order_list)
     global_order_list = Add_new_global_order(order1, global_order_list)
     global_order_list = Add_new_global_order(order2, global_order_list)
     global_order_list = Add_new_global_order(order3, global_order_list)
     global_order_list = Add_new_global_order(order4, global_order_list)

     // make my order list from the internal and global lists
     my_order_list = Make_my_order_list(internal_order_list, global_order_list)

     fmt.Println("This is my order list: ", my_order_list)

     internal_order_list = Delete_internal_order(internal_order_list)
     global_order_list = Delete_global_order(global_order_list)

     my_order_list = Make_my_order_list(internal_order_list, global_order_list)
     fmt.Println("This is my new order list: ", my_order_list)
}
*/

/* -------------------------------------------
   fmt.Println("detekterer et knappetrykk: ", button_pressed)
   new_order_floor := button_pressed.Floor
   new_order_button := button_pressed.Button

   if button_pressed.Button == global.BUTTON_COMMAND {
     for i := 0; i < global.NUM_INTERNAL_ORDERS; i++ {

     if internal_order_list[i].Order_state == Inactive {
       internal_order_list[i] = button_pressed
       fmt.Println("New internal order was added!")
       fmt.Println(internal_order_list)
       internal_order_list_chan <- internal_order_list
       break
     }
     if internal_order_list[i].Floor == new_order_floor && internal_order_list[i].Button == new_order_button {
       fmt.Println("The order is already in the internal order list.")
       break
     }
     }
   }

   if (button_pressed.Button == global.BUTTON_UP || button_pressed.Button == global.BUTTON_DOWN){
   for i := 0; i < global.NUM_GLOBAL_ORDERS; i++ {

     if global_order_list[i].Order_state == Inactive {
       global_order_list[i] = button_pressed
       fmt.Println("New external order was added!")
       fmt.Println(global_order_list)
       global_order_list_chan <- global_order_list
       break
     }
     if global_order_list[i].Floor == new_order_floor && global_order_list[i].Button == new_order_button {
       fmt.Println("The order is already in the global order list.")
       fmt.Println(global_order_list)
       break
     }
   }
 } -----------------------------------------------*/
