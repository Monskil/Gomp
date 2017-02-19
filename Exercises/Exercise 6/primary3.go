package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"encoding/binary"
)

func check_for_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func spawn_backup() {
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup2.go")
	err := cmd.Run()
	check_for_error(err)
}

func primary(start_number int){
	spawn_backup()

	// set up send-socket
	local_addr, _ := net.ResolveUDPAddr("udp", "")
	remote_addr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:20010")
	
	socket_send, err := net.DialUDP("udp", local_addr, remote_addr)
	check_for_error(err)

	// closing sockets
	defer socket_send.Close()

	msg := make([]byte, 1)
	for i := start_number; ; i++{
		fmt.Println(i)
		msg[0] = byte(i)
		socket_send.Write(msg)
		time.Sleep(100*time.Millisecond)
	}
}

func main(){
	primary(0)
}