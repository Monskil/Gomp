package main

import(
	"net"
	"log"
	"fmt"
	
)

func check_for_error(err error){
	if err != nil{
		log.Fatal(err)
	}
}


func main(){

	localAddr, err := net.ResolveTCPAddr("tcp", ":20010")
	remoteAddr, err := net.ResolveTCPAddr("tcp", "129.241.187.43:34933")
	check_for_error(err)

	socket, err := net.DialTCP("tcp", localAddr, remoteAddr)
	check_for_error(err)

	//make listener
	socket_listen, err := net.AcceptTCP();

	defer socket_listen.Close()

	for{
		socket.Write([1024]byte("Connect to: "))

		buffer := make([]byte, 1024)
		_, err := socket.Read(buffer)
		check_for_error(err)
		fmt.Println(buffer[:])

	}

}