package main

import (
	"fmt"
	"github.com/sabhiram/go-wol"
	"net"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")

	conn, err := net.DialUDP("udp", nil, udpAddr)

	defer conn.Close()

	macAddr := "D0-50-99-76-F3-32"

	mp, err := wol.New(macAddr)
	if err != nil {
		panic(err)
	}
	bs, err := mp.Marshal()
	if err != nil {
		panic(err)
	}

	n, err := conn.Write(bs)

	if err != nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
		return
	}
	fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
}
