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

func backup() int{
	fmt.Println("Hellooo, I'm the backup")

	listen_chan := make(chan int, 1)
	backup_value := 0
	go listen(listen_chan)

	for {
		select{
		case backup_value = <- listen_chan:
			time.Sleep(50*time.Millisecond)
			break
		case <- time.After(2*time.Second):
			fmt.Println("The primary is soo dead")
			return backup_value
		}
	}
}

func listen(listen_chan chan int){
	// set up listen socket
	port, _ := net.ResolveUDPAddr("udp", ":20010")
	socket_listen, err := net.ListenUDP("udp", port)
	check_for_error(err)

	// closing sockets
	defer socket_listen.Close()

	buffer := make([]byte, 1024)

	for {
		_, _, err := socket_listen.ReadFromUDP(buffer[:])
		check_for_error(err)

		listen_chan <- int(binary.LittleEndian.Uint64(buffer)) // convert  bytearray to int
		time.Sleep(100*time.Millisecond)
	}
}

func main(){
	backup_value := backup()
	primary(backup_value)
}