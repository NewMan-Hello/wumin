package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	//配置ip与端口
	server:=http.Server{
		Addr:"localhost:8080",
	}
	//http.HandleFunc("/url",Test)

	//直接调用匿名函数,查询fragment
	http.HandleFunc("/url", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Print(writer,request.URL.Fragment)
	})
	http.HandleFunc("/home",queryTest)
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

//func Test(w http.ResponseWriter,r *http.Request)  {
//	fmt.Print(w,r.URL.Fragment)
//}

//查询参数
func queryTest(w http.ResponseWriter,r *http.Request){
	query:=r.URL.Query()
	id:=query["id"]
	log.Println(id)
	name:=query.Get("name")
	log.Println(name)
}
