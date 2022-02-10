package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	addr := make(chan string)
	go startServer(addr)

	time.Sleep(time.Second * 3)

	data := []byte("[这里才是一个完整的数据包]")
	conn, err := net.DialTimeout("tcp", <-addr, time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}

// 服务端的控制台输出可以看出，存在三种类型的输出：
// 一种是正常的一个数据包输出。
// 一种是多个数据包“粘”在了一起，我们定义这种读到的包为粘包。
// 一种是一个数据包被“拆”开，形成一个破碎的包，我们定义这种包为半包。

// 为什么会出现半包和粘包？
// 客户端一段时间内发送包的速度太多，服务端没有全部处理完。于是数据就会积压起来，产生粘包。
// 定义的读的buffer不够大，而数据包太大或者由于粘包产生，服务端不能一次全部读完，产生半包。
// 什么时候需要考虑处理半包和粘包？
// TCP连接是长连接，即一次连接多次发送数据。
// 每次发送的数据是结构的，比如 JSON格式的数据 或者 数据包的协议是由我们自己定义的（包头部包含实际数据长度、协议魔数等）。
// 解决思路
// 定长分隔(每个数据包最大为该长度，不足时使用特殊字符填充) ，但是数据不足时会浪费传输资源
// 使用特定字符来分割数据包，但是若数据中含有分割字符则会出现Bug
// 在数据包中添加长度字段，弥补了以上两种思路的不足，推荐使用

// 通过上述分析，我们最好通过第三种思路来解决拆包粘包问题。
// Golang的bufio库中有为我们提供了Scanner，来解决这类分割数据的问题。
// Scanner为 读取数据 提供了方便的 接口。连续调用Scan方法会逐个得到文件的“tokens”，跳过 tokens 之间的字节。token 的规范由 SplitFunc 类型的函数定义。我们可以改为提供自定义拆分功能。
// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func exampleSplitFunc() {

	const input = "1234 5678 1234 1234 0987 8383 774"
	scanner := bufio.NewScanner(strings.NewReader(input))

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

}

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	addr <- l.Addr().String()
	for {
		// 监听到新的连接，创建新的 goroutine 交给 handleConn函数 处理
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")
	fmt.Println("new connection: ", conn.RemoteAddr())

	result := bytes.NewBuffer(nil)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read error:", err)
				break
			}
		} else {
			fmt.Println("received:", result.String())
		}
		result.Reset()
	}
}
