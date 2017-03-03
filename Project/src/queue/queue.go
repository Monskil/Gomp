package queue

import(
  "global"
  "driver"
)

//-- order states

type order struct{
  Button driver.button_t


}

//Global_order is the order queue the master is sending out, and is the order that should be acted out from
type Global_order struct{
		floor floor_t
		button button_t
		assigned_to int
    //order_state order_state
}

//When a slave receives a new order, they should send it to the master like this
type New_order struct{
			floor floor_t
			button button_t
			//elev_state elev_state
			elev_position int
			my_IP int
}

type my_orders struct {

}

var Global_order_matrix = []
