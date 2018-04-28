package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

type netInterface struct {
	iname string
	ipnet *net.IPNet
}

var intf []netInterface

//function to test binding to multiple interfaces.
//the trick is to use the default mask of the unique nic, and bind
//that with the netmask address to assure it only gets
//broadcast packets from that specific nic

// NOTE: unfortunately, if i only use the netmask, i can't get
// messages targetted for netmask > than that such as 255.255.255.255!!
func main() {

	localAddresses()
	port := 8001
	fmt.Println("Listening for UDP Broadcast Packet")

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP: net.IPv4(0, 0, 0, 0),
		//IP:   net.IPv4(192, 168, 1, 255),
		Port: port,
	})
	if err != nil {
		fmt.Println("Error listen: ", err)

	}

	//listen forever. it will pickup unicast and broadcast traffic
	for {
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("readfromudp: ", err)
		}
		fmt.Printf("Read %d bytes from: %v\n", read, remoteAddr)
		fmt.Print(hex.Dump(data[0:read]))

		iname := getInterfaceName(remoteAddr)
		fmt.Println("client was on interface: ", iname)
	}
}

//simple function to just collect all of theIPNet and interface names of each interface.
//filter out any that are not IPv4() since that is all we care about for broadcast
func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			_, ipNet, err := net.ParseCIDR(a.String())
			if err != nil {
				log.Printf("error parsing CIDR")
			}
			if ipNet.IP.To4() == nil || iface.Flags&net.FlagBroadcast == 0 || ipNet.IP.IsLoopback() {
				continue
			}

			tmpIntf := netInterface{iname: iface.Name, ipnet: ipNet}
			intf = append(intf, tmpIntf)
		}
	}
	fmt.Println(intf)
}

//Look through all the intf array of interfaces to find
//the match for the IP address to the IPNet mask
func getInterfaceName(raddr *net.UDPAddr) string {

	for _, a := range intf {
		if a.ipnet.Contains(raddr.IP) {
			return a.iname
		}
	}

	return "none"
}
