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

type master_msg struct {
	Address IP
	//global_liste
}

type slave_msg struct {
  Address IP
}
  
func Network_init(master bool) {
	//THIS IS NOW RUNNING AS SLAVE
  var receivePort, broadcastPort int
  if master {
    fmt.Println("Connecting as the master")
    receivePort = slavePort
    broadcastPort = masterPort


    master_sender := make(chan master_msg)
    master_receiver := make(chan slave_msg)
    go bcast.Transmitter(broadcastPort, master_sender)
    go bcast.Receiver(receivePort, master_receiver)
	 
    } else {
    fmt.Println("Connecting as a slave")
    receivePort = masterPort
    broadcastPort = slavePort
	  
   
    slave_sender := make(chan slave_msg)
    slave_receiver := make(chan master_msg)
    go bcast.Transmitter(broadcastPort, slave_sender)
    go bcast.Receiver(receivePort, slave_receiver)
    }

    var msg master_msg
    msg.IP = "Master sending msg"
    for {
    	master_sender <- master_msg
    	time.Sleep(1*time.Second)
    }

 }
