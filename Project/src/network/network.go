package network

import (
	"flag"
	"fmt"
	"network/bcast"
	"network/localip"
	"network/peers"
	"strconv"
	"time"
	"global"
	"queue"
)

const (
	master_port = 20079
	slave_port  = 20179
)

Local_ip, _ := localip.LocalIP()

type IP string

type Master_msg struct {
	Address IP
	Global_list [global.NUM_EXTERNAL_ORDERS]queue.Order
}

type Slave_msg struct {
	Address IP
	Internal_list [global.NUM_INTERNAL_ORDERS]queue.Order
	External_list [global.NUM_EXTERNAL_ORDERS]queue.Order
	Elevator_info queue.Elev_info
}

func Choose_master(ip_adresses peers.PeerUpdate) {
	highest_ip := 0
	str_local_ip, _ := Local_ip
	int_local_ip, _ := strconv.Atoi(str_local_ip[12:])
	for i := 0; i < len(ip_adresses.Peers); i++ {
		str_ip := ip_adresses.Peers[i]
		int_ip, _ := strconv.Atoi(str_ip[12:])
		if int_ip > highest_ip {
			highest_ip = int_ip
		}
	}
	if int_local_ip == highest_ip {
		fmt.Println("I am the master.")
		queue.Is_global = true
	} else {
		fmt.Println("I am a slave.")
		queue.Is_global = false
	}
}

func Network_handler() {
	var id string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.Parse()
	var new_info peers.PeerUpdate

	if id == "" {
		local_ip, err := Local_ip
		if err != nil {
			fmt.Println(err)
			local_ip = "Disconnected."
		}
		id = fmt.Sprintf(local_ip)

		peer_update_chan := make(chan peers.PeerUpdate)
		peer_transmit_enable_chan := make(chan bool)

		go peers.Transmitter(20243, id, peer_transmit_enable_chan)
		go peers.Receiver(20243, peer_update_chan)
		for {
			select {
			case new_info_catcher := <-peer_update_chan:
				new_info = new_info_catcher

				fmt.Printf("Peer update:\n")
				fmt.Printf("  Peers:    %q\n", new_info.Peers)
				fmt.Printf("  New:      %q\n", new_info.New)
				fmt.Printf("  Lost:     %q\n", new_info.Lost)
				Choose_master(new_info)
			}
		}
	}
}

func Network_setup(master bool) {
	var receive_port, broadcast_port int

	master_sender := make(chan Master_msg)
	master_receiver := make(chan Slave_msg)
	slave_sender := make(chan Slave_msg)
	slave_receiver := make(chan Master_msg)

	if master {
		fmt.Println("Connecting as the master.")
		receive_port = slave_port
		broadcast_port = master_port

		go bcast.Transmitter(broadcast_port, master_sender)
		go bcast.Receiver(receive_port, master_receiver)
		go master_transmit()

		for {
			select {
			case catch_msg_from_slave := <-master_receiver:
				fmt.Println("Master received : ", catch_msg_from_slave)
				queue.Master_msg_handler(catch_msg_from_slave)
			}
		}

	} else {
		fmt.Println("Connecting as a slave.")
		receive_port = master_port
		broadcast_port = slave_port

		go bcast.Transmitter(broadcast_port, slave_sender)
		go bcast.Receiver(receive_port, slave_receiver)
		go slave_transmit()

		for {
			select {
			case catch_msg_from_master := <-slave_receiver:
				fmt.Println("Slave received : ", catch_msg_from_master)
				queue.Slave_msg_handler(catch_msg_from_master)
			}
		}
	}
}

func master_transmit(){
	var msg Master_msg
	for {
		master_msg_to_send.Address = Local_ip
		master_msg_to_send.Global_list = queue.Global_order_list
		master_sender <- master_msg_to_send
		fmt.Println("Master sent the global list: ", master_msg_to_send.Global_list)
		time.Sleep(1 * time.Second)
	}
}

func slave_transmit(){
	var msg Slave_msg
	for {
		slave_msg_to_send.Address = Local_ip
		slave_msg_to_send.Internal_list = queue.Internal_order_list
		slave_msg_to_send.External_list = queue.External_order_list
		slave_msg_to_send.Elevator_info = queue.Elev_info
		slave_sender <- slave_msg_to_send
		fmt.Println("Slave sent the lists: ", )
		time.Sleep(1 * time.Second)
	}
}
