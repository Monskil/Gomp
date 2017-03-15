package msghandler

import (
  "network"
)

func Master_msg_handler(msg_from_slave network.Slave_msg){
  // Checks state on all orders for slave in external list, compares with global list
  // adds, delegates and updates state if order state is active,
  // updates state if order state is ready,
  // deletes order from global list if order state is finished
  slave_ip := msg_from_slave.Address
  internal_order_list := msg_from_slave.Internal_list
  external_order_list := msg_from_slave.External_list
  elev_info := msg_from_slave.Elevator_info

  for i := 0; i < global.NUM_EXTERNAL_ORDERS; i++ {
    if external_order_list[i].Assigned_to == slave_ip{
      if external_order_list[i].Order_state == queue.Active {
        queue.Add_new_global_order()
        //-- assign order
        queue.Global_list[j].Order_state == queue.Assigned
      } else{
        for j := 0; i < global.NUM_EXTERNAL_ORDERS; i++ {
          if external_order_list[i].Button == queue.Global_list[j].Button && external_order_list[i].Floor== queue.Global_list[j].Floor{
             if external_order_list[i].Order_state == queue.Ready {
              queue.Global_list[j].Order_state == queue.Ready
            } else if external_order_list[i].Order_state == queue.Finished{
              queue.Delete_global_order()
            }
          }
        }
      }
    }
  }
}

// if new order in external -> må legges til i global

func Slave_msg_handler(msg_from_master network.Master_msg){
  // Checks if any orders assigned to own ip in global, and then
  // checks if this order already exists in external list
  // if not -> add to external list
  my_ip := network.Local_ip
  master_ip := msg_from_master.Address
  global_order_list := msg_from_master.Global_list

  for i := 0; i < global.NUM_EXTERNAL_ORDERS; i++ {
    if global_order_list[i].Assigned_to == my_ip{
      for j := 0; j < global.NUM_EXTERNAL_ORDERS; j++ {
        if global_order_list[i].Button == queue.External_order_list[j].Button && global_order_list[i].Floor == queue.External_order_list[j].Floor{
          break
        }
      }
      global_order_list[i].Order_state = queue.Ready
      queue.Add_new_external_order(global_order_list[i])
      queue.Is_new_order = True
    }
  }
}
