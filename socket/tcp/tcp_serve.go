package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 处理完关闭连接
	defer func() {
		_ = conn.Close()
	}()

	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		// 针对当前连接做发送和接受操作
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)

		//conn.Write([]byte("ok")) // 发送数据

		// 将接受到的数据返回给客户端
		//_, err = conn.Write([]byte("ok"))
		//if err != nil {
		//	fmt.Printf("write from conn failed, err:%v\n", err)
		//	break
		//}
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
