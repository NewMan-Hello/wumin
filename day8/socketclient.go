package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main ()  {
	conn,err:=net.Dial("tcp","localhost:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}

	//客户端发送单行数据
	reader:=bufio.NewReader(os.Stdin)//os.stdin 代表标准输入

	//从终端读取用户输入，并准备发送给服务器
	line,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println("read err=",err)
	}

	//将line发送给服务器
	n,err:=conn.Write([]byte(line))
	if err!=nil{
		fmt.Println("conn.write err=",err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据",n)
}