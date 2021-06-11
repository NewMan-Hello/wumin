package main

import(
	"fmt"
	"strconv"
)

func main(){
	/*结果为A
	var a int=65
	b:=string(a)
   fmt.Print(b)*/

	var a int=65
	//输出字符"65"
	b:=strconv.Itoa(a)
	fmt.Println(b)
	//将字符转化为数字65
	a,_=strconv.Atoi(b)
	fmt.Print(a)


}