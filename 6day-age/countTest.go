package main
//计算机存储单位枚举
import "fmt"
const(
	B float64=1<<(iota*10)
	KB
	MB
	GB
)
func main(){
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

}