package main
import (
	"fmt"
	"runtime"
)
func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	c:=make( chan  bool)
	for i:=0;i<10 ;i++  {
		go Go(c,i)
	}
	for i:=0;i<10 ;i++  {
		<-c
	}
}
func Go(c chan bool,index int)  {
	a:=1
	for i:=0;i<1000000 ;i++  {
		a+=i
	}
	fmt.Println(index,a)
	c <-true
}