package Network

import (
  "fmt"
  )

const (
	masterPort = "20079"
	slavePort  = "20079"
)

type IP string

type master_msg struct {
	Address IP
	//global_liste
	Length  int
}

type slave_msg struct {
  Address IP
  
func Network_init(master bool) {
  var receivePort, broadcastPort string
  if master {
    fmt.Println("Connecting as the master")
    receivePort = masterPort
    broadcastPort = slavePort
	  

    master_sender := make(chan master_msg)
    master_receiver := make(chan slave_msg)
    go bcast.Transmitter(broadcastPort, master_sender)
    go bcast.Receiver(receivePort, master_receiver)
	 
    } else {
    fmt.Println("Connecting as a slave")
    receivePort = masterPort
    broadcastPort = slavePort
	  
   
    slave_sender := make(chan Slave_msg)
    slave_receiver := make(chan Master_msg)
    go bcast.Transmitter(broadcastPort, slave_sender)
    go bcast.Receiver(receivePort, slave_receiver)
    }
}

	
	
	
	
	//--------------------------------------------------------------------------------------------//
