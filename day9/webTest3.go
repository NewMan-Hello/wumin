package main
import(
	"io"
	"log"
	"net/http"
	"time"
)
var mux map[string]func(w http.ResponseWriter, r *http.Request)
func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}
	//路由设置
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello2
	mux["/bye"]=sayBye

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}
type myHandler struct {

}
func (*myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//判断路由是否与设定的一致
	if h,ok:=mux[r.URL.String()];ok{
		//执行匹配该路由的方法
		h(w,r)
		return
	}
	io.WriteString(w,"url:"+r.URL.String())
}
func  sayHello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"hello world!")

}
func  sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"bye world!")

}