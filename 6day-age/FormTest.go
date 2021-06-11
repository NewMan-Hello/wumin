package main
import (
	"fmt"
	"html/template"
	"net/http"
)
func main(){
	http.HandleFunc("/",index)
	http.HandleFunc("/doLogin",doLogin)
	http.ListenAndServe("location:8888",nil)
}
func index(w http.ResponseWriter,r *http.Request){
	t, _ :=template.ParseFiles("index.html")
	t.Execute(w,nil)
}
func doLogin(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if checkUsername(r.Form["username"][0]) {
		fmt.Print(w,"用户名",r.Form["username"][0],"密码",r.Form["password"][0])

	}else {
		t, _ :=template.ParseFiles("index.html")
		t.Execute(w,"用户名或密码错误")
	}
	
	
}
func checkUsername(username string) bool{
	if len(username)>=6 && len(username)<=16 {
		return  true
	}
	return false
}
