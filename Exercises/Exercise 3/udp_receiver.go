package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	portNum := "30000"

	service := ":" + portNum

	localAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.ListenUDP("udp", localAddr)

	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// write a message to serve


    /*message := []byte("Hello UDP server!")

	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}*/
	for {
		// receive message from server
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if(err != nil) {
			log.Fatal(err)
		}

		fmt.Println("UDP Server : ", addr)
		fmt.Println("Received from UDP server : ", string(buffer[:n]))
	}
}
