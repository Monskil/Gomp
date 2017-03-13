package network

import (
  "fmt"
  "network/bcast"
  "time"
  )

const (
	masterPort = 20079
	slavePort  = 20179
)

type IP string

type Master_msg struct {
	Address IP
	//global_liste
}

type Slave_msg struct {
  Address IP
}
  
func Network_init(master bool) {
	//THIS IS NOW RUNNING AS SLAVE
  var receivePort, broadcastPort int

   master_sender := make(chan Master_msg)
   master_receiver := make(chan Slave_msg)
   
    slave_sender := make(chan Slave_msg)
    slave_receiver := make(chan Master_msg)
  if master {
    fmt.Println("Connecting as the master")
    receivePort = slavePort
    broadcastPort = masterPort

    go bcast.Transmitter(broadcastPort, master_sender)
    go bcast.Receiver(receivePort, master_receiver)
	 
    } else {
    fmt.Println("Connecting as a slave")
    receivePort = masterPort
    broadcastPort = slavePort
	  
    go bcast.Transmitter(broadcastPort, slave_sender)
    go bcast.Receiver(receivePort, slave_receiver)
    }

    var msg Master_msg
    msg.Address = "Master sending msg"
    for {
    	master_sender <- msg
    	time.Sleep(1*time.Second)
    }

 }
