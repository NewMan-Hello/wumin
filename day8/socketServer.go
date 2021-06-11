package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	//循环接受客户端发送的数据
	defer conn.Close()

	for{
		//创建一个切片
		buf:=make([]byte,1024)
		//等待客户端通过conn发送信息，如果没有发送就会堵塞
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("服务器的read err=",err)
			return
		}
		//显示客户端发送的消息
		fmt.Print(buf[:n])
	}
}
func main(){
	fmt.Println("开始监听:")
	//tcp 网络协议
	//8888 表示端口号
	listen,err:=net.Listen("tcp","0.0.0.0:8888")
	if err!=nil {
		fmt.Println("listen err=",err)
		return
	}
	defer  listen.Close()//延时关闭listen
	//等待链接
	for{
		fmt.Println("等待客户端链接")
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("accept err=",err)
		}else {
			fmt.Printf("accpet suc conn=%v 客户端ip=%v\n",conn,conn.RemoteAddr().String())
		}
		go process(conn)
	}

}