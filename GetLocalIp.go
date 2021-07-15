package main

import (
	"log"
	"net"
	"strings"
)

func GetLocalIp() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	defer conn.Close()
	if err != nil {
		log.Printf("Get ip error: %v", err)
		return ""
	}
	ip = strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip
}
