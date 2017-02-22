package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	//	"encoding/binary"
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

func backup() int {
	fmt.Println("Hellooo, I'm the backup")

	// set up listen socket
	port, _ := net.ResolveUDPAddr("udp", ":20010")
	socket_listen, err := net.ListenUDP("udp", port)
	check_for_error(err)

	// closing sockets
	defer socket_listen.Close()

	listen_chan := make(chan int, 1)
	go_on := make(chan int, 1)
	primary_dead_chan := make(chan int, 0)
	primary_alive := true
	timer := time.NewTimer(2 * time.Second)

	backup_value := 0
	backup_value_old := 0
	backup_value_new := 0

	go func() {
		for {
			select {
			case backup_value_new = <-listen_chan:
				//fmt.Println("I got the value ", backup_value)
				backup_value_old = backup_value
				backup_value = backup_value_new
				timer.Reset(2 * time.Second)
				time.Sleep(50 * time.Millisecond)
				break
			case <-timer.C:
				fmt.Println("The primary is soo dead")
				fmt.Println("My startup value will be ", backup_value_old)
				primary_dead_chan <- backup_value
				fmt.Println("Channel okay")
				primary_alive = false
				fmt.Println("Primary now sat to dead")
				time.Sleep(10 * time.Second)
			}
		}
	}()

	buffer := make([]byte, 1024)
	to_listen := 0

	for {
		fmt.Println("In the for-loop bby")

		if primary_alive == true {
			go_on <- backup_value
		}

		//switch primary_alive {
		select {
		case <-primary_dead_chan:
			//case false:
			fmt.Println("I know it's dead")
			return backup_value_old
		case <-go_on:
			//case true:
			n, _, err := socket_listen.ReadFromUDP(buffer[:])
			check_for_error(err)

			//listen_chan <- int(binary.LittleEndian.Uint64(buffer)) // convert  bytearray to int
			to_listen, _ = strconv.Atoi(string(buffer[:n]))
			listen_chan <- to_listen
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("At the end of the for-loop sweetheart")
	}
}

func main() {
	backup_value := backup()
	//fmt.Println("My start value is now ", backup_value)
	primary(backup_value)
}
