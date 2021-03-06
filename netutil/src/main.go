package main

import (
        "flag"
	"fmt"
	"net"
	"time"
)

func main() {
	ipPtr := flag.String("ip", "127.0.0.1", "server ip")
	portPtr := flag.String("port", "5000", "server port")
	flag.Parse()

	res, err := sendUDP(*ipPtr + ":" + *portPtr, "hello")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
func sendUDP(addr, msg string) (string, error) {

	conn, _ := net.Dial("udp", addr)

	// send to socket
	_, err := conn.Write([]byte(msg))

	// listen for reply
	bs := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(3 * time.Second))
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}
