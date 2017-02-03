package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func check_for_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//TCP setup
	remoteAddr, err := net.ResolveTCPAddr("tcp", "129.241.187.43:34933")
	check_for_error(err)

	socket, err := net.DialTCP("tcp", nil, remoteAddr)
	check_for_error(err)

	//Make listener
	localAddr, err := net.ResolveTCPAddr("tcp", "129.241.187.43:34933")
	check_for_error(err)
	socket_listen, err := net.ListenTCP("tcp", localAddr)
	check_for_error(err)

	//server connect back
	_, err = socket.Write([]byte("Connect to: 129.241.187.158"))
	check_for_error(err)

	socket_connect, err := socket_listen.AcceptTCP()
	check_for_error(err)
	message := "hellohello"

	for {

		var buffer [1024]byte
		_, err := socket_connect.Read((buffer[:]))
		check_for_error(err)
		fmt.Println(buffer[:])

		_, err = socket_connect.Write([]byte(message))
		check_for_error(err)

		time.Sleep(2 * time.Second)
	}

	/*
		buffer := make([]byte, 1024)
		n, err := socket.Read(buffer)
		fmt.Println(string(buffer[:n]))
	*/

	/*
	   for{
	   	socket.Write([1024]byte("Connect to: 129.241.187.158\0"))

	   	buffer := make([]byte, 1024)
	   	_, err := socket.Read(buffer)
	   	check_for_error(err)
	   	fmt.Println(buffer[:])

	   }
	*/
}
