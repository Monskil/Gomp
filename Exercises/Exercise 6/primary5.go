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
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup5.go")

	err := cmd.Run()
	check_for_error(err)
}

func primary(start_number int) {
	// spawn backup
	spawn_backup()

	// set up send-socket
	local_addr, _ := net.ResolveUDPAddr("udp", "")
	remote_addr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:20010")
	socket_send, err := net.DialUDP("udp", local_addr, remote_addr)
	check_for_error(err)

	// closing sockets
	defer socket_send.Close()

	number := start_number
	str_number := ""

	go func() {
		for {
			fmt.Println(number)
			time.Sleep(1 * time.Second)
			number += 1
		}
	}()

	for {
		// send message
		str_number = strconv.Itoa(number)
		socket_send.Write([]byte(str_number))
		time.Sleep(10 * time.Millisecond)
	}

}

func main() {
	primary(0)
}
