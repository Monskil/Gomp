package main

// bare copy-pasta, ikke i orden i de thele tatt
import (
	"fmt"
	"log"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP  client : ", addr)
	fmt.Println("Received from UDP client : ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	hostName := "localhost"
	portNum := "30000"
	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", service)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server up and listening on port 30000")
	defer ln.Close()

	for {
		//wait for UDP client to connect
		handleUDPConnection(ln)
	}

}
