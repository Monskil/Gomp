package testq

import(
  //def "global"
  //"driver"
  "fmt"
)


// ---- global -----------------------------------------------
type button_t int

const (
	BUTTON_UP = iota
	BUTTON_DOWN
	BUTTON_COMMAND
)

type floor_t int

const (
	FLOOR_1 = iota
	FLOOR_2
	FLOOR_3
	FLOOR_4
)

type motor_direction_t int

const (
	DIR_DOWN = -1 << iota
	DIR_STOP
	DIR_UP
)

type assigned_t int

const (
  NONE = iota
  ELEV_1
  ELEV_2
  ELEV_3
)

const NUM_GLOBAL_ORDERS = 6
const NUM_INTERNAL_ORDERS = 4
const NUM_ORDERS = NUM_GLOBAL_ORDERS + NUM_INTERNAL_ORDERS
//---------------------------------------------------------


// order states
const (
  inactive = iota
  added
  assigned
  ready
  active
)

type order struct{
  button button_t
  floor floor_t
  order_state int
  assigned_to assigned_t
}

type elev_info struct{
  elev_ip int
  elev_last_floor floor_t
  elev_dir motor_direction_t
  elev_state int
}

// declare variables
var order_state int

func make_my_order_list(internal_order_list [NUM_INTERNAL_ORDERS]order, global_order_list [NUM_GLOBAL_ORDERS]order)[NUM_ORDERS]order{
  // add first the elements from the internal list
  // then the elements of the global list

  var my_order_list [NUM_ORDERS] order
  for i := 0; i < NUM_ORDERS; i++{
      if i < NUM_INTERNAL_ORDERS {
        my_order_list[i] = internal_order_list[i]
      } else{
        my_order_list[i] = global_order_list[i - NUM_INTERNAL_ORDERS]
      }
  }
  return my_order_list
}

func make_new_order(button button_t, floor floor_t, order_state int, assigned_to assigned_t) order{
  var new_order order

  new_order.button = button
  new_order.floor = floor
  new_order.order_state = order_state
  new_order.assigned_to = assigned_to

  return new_order
}

func add_new_internal_order(new_order order, internal_order_list [NUM_INTERNAL_ORDERS]order)[NUM_INTERNAL_ORDERS]order{
  new_order_floor := new_order.floor

  for i := 0; i < NUM_INTERNAL_ORDERS; i++{
    if internal_order_list[i].order_state == inactive {
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

func add_new_global_order(new_order order, global_order_list [NUM_GLOBAL_ORDERS]order)[NUM_GLOBAL_ORDERS]order{
  new_order_floor := new_order.floor
  new_order_button := new_order.button

  for i := 0; i < NUM_GLOBAL_ORDERS; i++{
    if global_order_list[i].order_state == inactive {
      global_order_list[i] = new_order
      fmt.Println("New order was added!")
      return global_order_list
    }
    if global_order_list[i].floor == new_order_floor && global_order_list[i].button == new_order_button{
      fmt.Println("The order is already in the internal order list.")
      return global_order_list
    }
  }
  fmt.Println("Something went wrong, we didn't add any new orders.")
  return global_order_list
}

// -- når man sletter en bestilling så må man flytte alle de andre bestillingene et hakk bort

// initial values
func Init_queue(){
  fmt.Println("Helloo, I'm initializing the queue, yayy!")

  // make lists
  var internal_order_list [NUM_INTERNAL_ORDERS] order
  var global_order_list [NUM_GLOBAL_ORDERS] order
  var my_order_list [NUM_ORDERS] order

  // test example making orders
  var order1 = make_new_order(BUTTON_UP, FLOOR_2, ready, ELEV_2)
  var order2 = make_new_order(BUTTON_COMMAND, FLOOR_1, active, ELEV_3)
  var order3 = make_new_order(BUTTON_UP, FLOOR_1, active, ELEV_1)
  var order4 = make_new_order(BUTTON_COMMAND, FLOOR_2, active, ELEV_2)

  //global_order_list[0] = order1
  //global_order_list[5] = order1
  //internal_order_list[3] = order2

  internal_order_list = add_new_internal_order(order2, internal_order_list)
  internal_order_list = add_new_internal_order(order1, internal_order_list)
  internal_order_list = add_new_internal_order(order2, internal_order_list)

  global_order_list = add_new_global_order(order2, global_order_list)
  global_order_list = add_new_global_order(order1, global_order_list)
  global_order_list = add_new_global_order(order2, global_order_list)
  global_order_list = add_new_global_order(order3, global_order_list)
  global_order_list = add_new_global_order(order4, global_order_list)

  // make my order list from the internal and global lists
  my_order_list = make_my_order_list(internal_order_list, global_order_list)

  fmt.Println("This is my global order list: ", global_order_list)
  fmt.Println("This is my order list: ", my_order_list)
}
