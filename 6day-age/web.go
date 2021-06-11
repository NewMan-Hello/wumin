package main

import (
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/logout",logout)
	http.ListenAndServe("location:8888",nil)
}
func handler(rw http.ResponseWriter,req *http.Request){
	req.ParseForm()
	if len(req.Form["name"])>0 {
		fmt.Fprint(rw,"hello",req.Form["name"][0])
	}else {
		fmt.Fprint(rw,"helloworld")
	}

}
func logout(rw http.ResponseWriter,req *http.Request ){
	fmt.Fprint(rw,"logout")
}
