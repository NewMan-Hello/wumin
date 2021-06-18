package main
import(
	"fmt"
	"net/http"
)
func writeExample(w http.ResponseWriter,r *http.Request){
	str:="hello word"
	//write将接收的数据写入相应的http的body中
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter,r *http.Request){
	//通过writeHeader来发送HTTP状态码
	w.WriteHeader(501)
	fmt.Println("NO such service,try again")
}

func header(w http.ResponseWriter,r *http.Request){
	//通过header.set方法来修改headers的内容，并返回给客户端
	w.Header().Set("location","http://google.com")
	w.WriteHeader(302)
}
func main(){
	server:=http.Server{
		Addr:"localhost:8080",
	}
	http.HandleFunc("/write",writeExample)
	http.HandleFunc("/writeHeader",writeHeaderExample)
	http.HandleFunc("/header",header)
	server.ListenAndServe()
}