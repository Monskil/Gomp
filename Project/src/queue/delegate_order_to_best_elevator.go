package queue

//tar inn ordern, best suited_elevator ++
//sjekker om orderen ligger i noen av de ulike heisenes lister, dersom den gjør det kan ordren slettes for da blir den utført
//dersom den ikke ligger inne må vi delegere ordren til heisen best_suited elevator
//returnerer heisen som får den?


func delegate_order_to_best_elevator (elevator [global.NUM_ELEV]Elev_info, num_elevators_online int, new_order Order, external_order_list [global.NUM_GLOBAL_ORDERS]Order) Assigned_t{
  var assigned_elevator Assigned_t
  my_order := new_order
  compare_cost(elevator, num_elevators_online) //best elevator
  
  Add_new_external_order(new_order, external_order_list)  //add
return assigned elevator
}

