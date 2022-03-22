package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	defer conn.Close()
	if err != nil {
		fmt.Println("dial err", err)
		return
	}
	data := "hello"
	n, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Write err", err)
		return
	}
	fmt.Println("send bytes", n)

}