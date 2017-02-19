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
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup.go")

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

func backup() {
	// set up listen socket
	fmt.Println("Hellooo")
	value := 0
	port, _ := net.ResolveUDPAddr("udp", ":20010")
	socket_listen, err := net.ListenUDP("udp", port)
	check_for_error(err)

	// closing sockets
	defer socket_listen.Close()

	timer := time.NewTimer(time.Second * 2)
	primary_alive := true
	go func() {
		<-timer.C
		primary_alive = false
		fmt.Println("Primary is sooo dead.")
		primary(value)
	}()

	for {
		buffer := make([]byte, 1024)

		_, _, err := socket_listen.ReadFromUDP(buffer[:])
		check_for_error(err)

		x, _ := strconv.Atoi(string(buffer))
		value = x
		timer.Reset(time.Second * 2)
		if primary_alive == false {
			break
		}

	}
}

func main() {
	backup()
}
