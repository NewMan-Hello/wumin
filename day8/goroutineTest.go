package main

import "fmt"
var c chan string
func talk(){
	i:=0
	for{
		//获取c中的值，没有就进行等待
		fmt.Println(<-c)
		//向c传入值
		c<-fmt.Sprintf("from talk: hi,#%d",i)
		i++
	}
}
func main()  {
	c=make(chan string)
	go talk()
	for i:=0;i<10 ;i++  {
		//存入值
		c<-fmt.Sprintf("from main: hello,#%d",i)
		//获取值，无则等待
		fmt.Println(<-c)
	}
}