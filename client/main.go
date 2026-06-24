package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	// 连接服务器
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("net.Dial err: %v\n", err)
		return
	}

	// 协程接收服务端消息
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("server closed")
					return
				}
				fmt.Printf("conn.Read err: %v\n", err)
			}
			// 客户端打印服务端消息
			fmt.Print(string(buf[:n]))
		}
	}()
	// 接收终端消息
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		trimStr := strings.Trim(str, "\n")
		if trimStr == "exit" {
			return
		}
		if err != nil {
			fmt.Printf("recv terminal msg err: %v\n", err)
			continue
		}
		// 发送给服务端
		conn.Write([]byte(str))
	}
}
