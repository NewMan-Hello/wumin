package main
import "fmt"

type Num int
func main(){
 var num Num
 num.Increase(100)
 fmt.Print(num)
}
func (num *Num) Increase(n int){
	*num+=Num(n)
}
