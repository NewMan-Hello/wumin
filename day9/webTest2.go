package main

import (
	"io"
	"log"
	"net/http"
)
func main(){
	//serveMux http请求路由器
	mux:=http.NewServeMux()
	//设置路由
	mux.Handle("/",&myHandle{})
	mux.HandleFunc("/hello",sayHello1)
	err:=http.ListenAndServe(":8080",mux)
	if err!=nil{
		log.Fatal(err)
	}
}

type myHandle struct {

}
func (*myHandle) ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"url:"+r.URL.String())
}
func  sayHello1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"hello world!")

}
