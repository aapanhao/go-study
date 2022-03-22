// 服务端
//

package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	fmt.Println("获取连接", conn)
	fmt.Println("获取连接，地址是", conn.RemoteAddr())

	str := make([]byte, 1024)
	n, err := conn.Read(str)
	if err != nil {
		fmt.Println("Read err, ", err)
		return
	}
	fmt.Println("数据长度", n, string(str))
}

func main() {
	ln, err := net.Listen("tcp", ":8088")
	defer ln.Close()
	if err != nil {
		fmt.Println("listen err, ", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept err, ", err)
			return
		}
		go process(conn)
	}
}
