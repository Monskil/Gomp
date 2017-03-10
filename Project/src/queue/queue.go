// The orders have 6 states:
// Inactive: no order
// Added: the order has been added to the list, but not yet assigned
// Assigned: the order has been assigned
// Ready: verified that order is known at assigned elevator, waiting in line to be executed
// Active: the order is being executed
// Finished: the order is finished

package queue

import(
  "global"
  //"driver"
  "fmt"
)

// order states
const (
  Inactive = iota
  Added
  Assigned
  Ready
  Active
  Finished
)

type Order struct{
  Button global.Button_t
  Floor global.Floor_t
  Order_state int
  Assigned_to global.Assigned_t
}

type Elev_info struct{
  Elev_ip int
  Elev_last_floor global.Floor_t
  Elev_dir global.Motor_direction_t
  Elev_state int
}

func Make_my_order_list(internal_order_list [global.NUM_INTERNAL_ORDERS]Order, global_order_list [global.NUM_GLOBAL_ORDERS]Order)[global.NUM_ORDERS]Order{
  // add first the elements from the internal list
  // then the elements of the global list

  var My_order_list [global.NUM_ORDERS] Order
  for i := 0; i < global.NUM_ORDERS; i++{
      if i < global.NUM_INTERNAL_ORDERS {
        my_order_list[i] = internal_order_list[i]
      } else{
        my_order_list[i] = global_order_list[i - global.NUM_INTERNAL_ORDERS]
      }
  }
  return my_order_list
}

func Make_new_order(button global.Button_t, floor global.Floor_t, order_state int, assigned_to global.Assigned_t) Order{
  // make new order with the elements you choose
  var new_order Order

  new_order.button = button
  new_order.floor = floor
  new_order.order_state = order_state
  new_order.assigned_to = assigned_to

  return new_order
}

func Add_new_internal_order(new_order Order, internal_order_list [global.NUM_INTERNAL_ORDERS]Order)[global.NUM_INTERNAL_ORDERS]Order{
  // add new order in the internal list, if it's not already there
  new_order_floor := new_order.floor

  for i := 0; i < global.NUM_INTERNAL_ORDERS; i++{
    if internal_order_list[i].order_state == Inactive {
      internal_order_list[i] = new_order
      fmt.Println("New order was added!")
      return internal_order_list
    }
    if internal_order_list[i].floor == new_order_floor{
      fmt.Println("The order is already in the internal order list.")
      return internal_order_list
    }
  }
  fmt.Println("Something went wrong, we didn't add any new orders.")
  return internal_order_list
}

func Add_new_global_order(new_order Order, global_order_list [global.NUM_GLOBAL_ORDERS]Order)[global.NUM_GLOBAL_ORDERS]Order{
  // add new order in the global list, if it's not already there
  new_order_floor := new_order.floor
  new_order_button := new_order.button

  for i := 0; i < global.NUM_GLOBAL_ORDERS; i++{
    if global_order_list[i].order_state == Inactive {
      global_order_list[i] = new_order
      fmt.Println("New order was added!")
      return global_order_list
    }
    if global_order_list[i].floor == new_order_floor && global_order_list[i].button == new_order_button{
      fmt.Println("The order is already in the global order list.")
      return global_order_list
    }
  }
  fmt.Println("Something went wrong, we didn't add any new orders.")
  return global_order_list
}

func Delete_internal_order(internal_order_list [global.NUM_INTERNAL_ORDERS]Order) [global.NUM_INTERNAL_ORDERS]Order{
  // delete all finished orders in the internal list by moving all the "later" orders one step forward
  // make the last element in the list an "empty" order
  clean_order := make_new_order(global.BUTTON_UP, global.FLOOR_1, Inactive, global.NONE)

  for i := 0; i < global.NUM_INTERNAL_ORDERS; i++{
    if internal_order_list[i].order_state == Finished {
      fmt.Println("An internal order is marked finished.")
      for j:= i; j < global.NUM_INTERNAL_ORDERS; j++{
        if j < global.NUM_INTERNAL_ORDERS-1{
          fmt.Println("Moving order.")
          internal_order_list[j] = internal_order_list[j+1]
        } else if j == global.NUM_INTERNAL_ORDERS -1{
          fmt.Println("Adding last clean order.")
          internal_order_list[j] = clean_order
        }
      }
    }
  }
  return internal_order_list
}

func Delete_global_order(global_order_list [global.NUM_GLOBAL_ORDERS]Order) [global.NUM_GLOBAL_ORDERS]Order{
  // delete all finished orders in the global list by moving all the "later" orders one step forward
  // make the last element in the list an "empty" order
  clean_order := make_new_order(global.BUTTON_UP, global.FLOOR_1, Inactive, global.NONE)

  for i := 0; i < global.NUM_GLOBAL_ORDERS; i++{
    if global_order_list[i].order_state == Finished {
      fmt.Println("A global order is marked finished.")
      for j:= i; j < global.NUM_ORDERS; j++{
        if j < global.NUM_GLOBAL_ORDERS-1{
          fmt.Println("Moving order.")
          global_order_list[j] = global_order_list[j+1]
        } else if j == global.NUM_GLOBAL_ORDERS -1{
          fmt.Println("Adding last clean order.")
          global_order_list[j] = clean_order
        }
      }
    }
  }
  return global_order_list
}

// initial values
// and also some testing
func Init_queue(){
  fmt.Println("Helloo, I'm initializing the queue, yayy!")

  // make lists
  var internal_order_list [global.NUM_INTERNAL_ORDERS] Order
  var global_order_list [global.NUM_GLOBAL_ORDERS] Order
  var my_order_list [global.NUM_ORDERS] Order

  // test example making orders
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
