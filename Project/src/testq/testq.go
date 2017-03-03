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

  return internal_order_list
}

func add_new_global_order(new_order order, global_order_list [NUM_INTERNAL_ORDERS]order)[NUM_INTERNAL_ORDERS]order{

  return global_order_list
}

// initial values
func Init_queue(){
  fmt.Println("Helloo, I'm initializing the queue, yayy!")

  // make lists
  var internal_order_list [NUM_INTERNAL_ORDERS] order
  var global_order_list [NUM_GLOBAL_ORDERS] order
  var my_order_list [NUM_ORDERS] order

  // test example making an order
  var order1 order
  order1.button = BUTTON_UP
  order1.floor = FLOOR_3
  order1.order_state = inactive
  order1.assigned_to = NONE

  var order2 = make_new_order(BUTTON_COMMAND, FLOOR_1, active, ELEV_3)

  global_order_list[0] = order1
  global_order_list[5] = order1
  internal_order_list[3] = order2

  // make my order list from the internal and global lists
  my_order_list = make_my_order_list(internal_order_list, global_order_list)

  fmt.Println("This is my global order list: ", global_order_list)
  fmt.Println("This is my order list: ", my_order_list)
}
