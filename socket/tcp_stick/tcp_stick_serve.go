package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"pocket_lab/socket/helper"
)

func process(conn net.Conn) {
	// 处理完关闭连接
	defer func() {
		_ = conn.Close()
	}()

	reader := bufio.NewReader(conn)
	// 针对当前连接做发送和接受操作
	for {
		msg, err := helper.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	defer listen.Close()
	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
}
