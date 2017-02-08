package main

import (
 	"fmt"
	"log"
	"net"
	"time"
)


func check_for_error(err error){
	if err != nil{
		log.Fatal(err)
	}
}

//Get your own IP, sånn at alle tre pc'ene kan gjøre det.

//Sverres get_your_own_IP
func get_your_own_IP()(string, error){
	if localIP == "" {
		conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
		if err != nil {
			return "", err
		}
		defer conn.Close()
		localIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return localIP, nil	
}

func set_up_broadcast_socket_send(){
}

func set_up_broadcast_socket_listen(){
}
