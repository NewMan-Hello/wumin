package main
import(
	"encoding/json"
	"log"
	"net/http"
)
type Company struct {
	ID    int   `json:"id"`
	Name   string 	`json:"name"`
	Country string `json:"country"`
}
func main(){
	server:=http.Server{
		Addr:"localhost:8080",
	}
	http.HandleFunc("/companies", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			//读取json数据
			//创建一个解码器
			dec:=json.NewDecoder(request.Body)
			//定义一个结构体接收
			company:=Company{}
			//解码
			err:=dec.Decode(&company)
			if err != nil{ //解码失败，则抛出错误，并修改状态码
				log.Println(err.Error())
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			//如果解码成功，则转回为json格式返回给客户端
			//创建一个编码器
			enc:=json.NewEncoder(writer)
			//编码
			err=enc.Encode(company)
			if err!=nil{
				log.Println(err.Error())
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			writer.WriteHeader(http.StatusInternalServerError)
		}
		server.ListenAndServe()
	})
}

