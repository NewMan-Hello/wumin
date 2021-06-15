package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
func process(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(1024)
	fmt.Print(w,r.PostForm)
	//获取文件
	fileHeader:=r.MultipartForm.File["uploaded"][0]
	//打开文件
	file,err:=fileHeader.Open()
	if err!=nil{
		//读取文件,转换成切片
		data,err:=ioutil.ReadAll(file)
		if err!=nil{
			//输出文件
			fmt.Fprint(w,string(data))
		}
	}
}
func main(){
	server:=http.Server{
		Addr:              "localhost:8080",
	}
	http.HandleFunc("/process",process)

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}