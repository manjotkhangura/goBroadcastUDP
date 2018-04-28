package main

import (
	//"bitbucket.org/smartcast/conf"
	"log"
	"net"
)

//exercise multiple ways of sending a broadcast request
//used for testing server.go broadcast listener
func main() {

	message := []byte("hello")
	ipBcast := net.IPv4(192, 168, 1, 255)
	//ipBcast := net.IPv4(255, 255, 255, 255)

	//set local ip binding
	ipLocal := net.IPv4(192, 168, 1, 10)
	//ipLocal := net.IPv4(10, 10, 10, 129)
	//ipLocal := net.IPv4(127, 0, 0, 1)
	l := net.UDPAddr{IP: ipLocal, Port: 8002}

	socket, err := net.DialUDP("udp4", &l, &net.UDPAddr{
		IP:   ipBcast,
		Port: 8001,
	})
	if err != nil {
		log.Printf("Error while starting TXQueue: %s, \n%v\n %v\n", err, ipBcast, ipLocal)
		return
	}

	socket.Write(message)
}
