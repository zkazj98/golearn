package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("n:", n)
		recvStr := string(buf[:n])
		fmt.Printf("收到客户端消息:", recvStr)
		conn.Write([]byte(recvStr))
	}
}
