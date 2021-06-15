package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	server:=http.Server{
		Addr:              "localhost:8080",
		//Handler:           nil,
		//TLSConfig:         nil,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	//获取请求头信息
	http.HandleFunc("/header", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer,request.Header)
		//返回数组类型
		fmt.Println(writer,request.Header["Accept-Encoding"])
		//返回字符串类型
		fmt.Println(writer,request.Header.Get("Accept-Encoding"))
	})

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}