package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"time"
)

func check_for_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func spawn_backup() {
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup4.go")
	err := cmd.Run()
	check_for_error(err)
}

func primary(start_number int) {
	spawn_backup()

	// set up send-socket
	local_addr, _ := net.ResolveUDPAddr("udp", "")
	remote_addr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:20010")

	socket_send, err := net.DialUDP("udp", local_addr, remote_addr)
	check_for_error(err)

	// closing sockets
	defer socket_send.Close()

	number := start_number

	//msg := make([]byte, 1)
	//for number := start_number; ; number++ {
	for {
		fmt.Println(number)
		//msg[0] = byte(i)
		str_number := strconv.Itoa(number)
		socket_send.Write([]byte(str_number))
		time.Sleep(100 * time.Millisecond)

		number++
	}
}

func main() {
	primary(0)
}
