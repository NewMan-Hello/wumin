package main
import(
	"html/template"
	"net/http"
)
func main(){
	server:=http.Server{
		Addr:"localhost:8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}
func process(w http.ResponseWriter,r *http.Request){
	t,_:=template.ParseFiles("t1.html")
	//Execute适用于执行单模板，存在多模版时只执行第一个
	t.Execute(w,"hello world")

	ts,_:=template.ParseFiles("t1.html","t2.html")
	//ExecuteTemplate在多模板时可以指定模板
	ts.ExecuteTemplate(w,"t2.html","hello world")
}
