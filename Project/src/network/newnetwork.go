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
    } else {
    fmt.Println("Connecting as a slave")
    receivePort = masterPort
    broadcastPort = slavePort
    }
  
  
func Network_init(master bool, sendChannel, receiveChannel chan UDPMessage, networkLogger log.Logger) {
	var localPort, broadcastPort string
	if master {
		networkLogger.Print("Connecting as master")
		localPort = masterPort
		broadcastPort = slavePort
	} else {
		networkLogger.Print("Connecting as slave")
		localPort = slavePort
		broadcastPort = masterPort
	}

	}
