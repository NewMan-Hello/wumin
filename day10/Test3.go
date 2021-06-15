package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	server:=http.Server{
		Addr:              "localhost:8080",
	}
	//获取消息体
	http.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
		len:=request.ContentLength
		//通过body来接收数据
		body:=make([]byte,len)
		request.Body.Read(body)
		fmt.Print(writer,string(body))
		})

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}