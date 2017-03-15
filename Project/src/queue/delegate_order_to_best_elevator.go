package queue


func delegate_order_to_best_elevator(elevator [global.NUM_ELEV]Elev_info, num_elevators_online int, new_order Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order)Assigned_t {
  var assigned_elevator Assigned_t
  
  compare_cost(elevator, num_elevators_online) 
  
  assigned_elevator = Elev_info.Elev_ip
  Add_new_global_order(new_order, external_order_list) 
  
  return assigned elevator
}

