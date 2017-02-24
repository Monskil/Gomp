
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

func spawn_primary() {
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run primary.go " + value)

	err := cmd.Run()
	check_for_error(err)
}


func backup() {
	// set up listen socket
	fmt.Println("Hellooo, I'm the backup")
	value := 0
	port, _ := net.ResolveUDPAddr("udp", ":20010")
	socket_listen, err := net.ListenUDP("udp", port)
	check_for_error(err)
	fmt.Println("Have set up the socket")

	// closing sockets
	defer socket_listen.Close()

	timer := time.NewTimer(2 * time.Second)
	primary_alive := true
	go func() {
		<-timer.C
		fmt.Println("Timers out")
		primary_alive = false
		return
	}()

	go func() {
		for{

		buffer := make([]byte, 1024)

		_, _, err := socket_listen.ReadFromUDP(buffer[:])
		check_for_error(err)

		x, _ := strconv.Atoi(string(buffer))
		value = x
		

		timer.Reset(2 * time.Second)
		}

	}()

	for {
		fmt.Printf("%t", primary_alive)
		if primary_alive == false {

			fmt.Println("Primary is sooo dead.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("I'm now spawning the primary with the value", value)
			spawn_primary()
			return
		}

	}

}

func main() {
	backup()
}
