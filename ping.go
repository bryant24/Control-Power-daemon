package main

import (
	"fmt"
	"net"
	"time"
)

func Ping(ip string) bool {
	recvBuf1 := make([]byte, 2048)
	payload := []byte{0x08, 0x00, 0x4d, 0x4b, 0x00, 0x01, 0x00, 0x10, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69}
	Time, _ := time.ParseDuration("1s")
	conn, err := net.DialTimeout("ip4:icmp", ip, Time)
	if err != nil {
		fmt.Println("bibi")
		return false
	}
	_, err = conn.Write(payload)
	if err != nil {
		return false
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	num, err := conn.Read(recvBuf1[0:])
	conn.SetReadDeadline(time.Time{})
	if string(recvBuf1[0:num]) != "" {
		return true
	}

	return false
}
