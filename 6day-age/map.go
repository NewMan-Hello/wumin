package main

import (
	"fmt"
)
func main(){
	var m map[int]string
	m=make(map[int]string)
	fmt.Println(m)
	m[1]="ok"
	fmt.Println(m[1])

	//多级map
	m1:=make(map[int]map[int]string)
	//ok为Boolean，默认为true
	a,ok:=m1[2][1]
	fmt.Println(a,ok)
	if !ok{
		m1[2]=make(map[int]string)
	}
	m1[2][1]="good"
	a,ok=m1[2][1]
	fmt.Println(a,ok)
}

