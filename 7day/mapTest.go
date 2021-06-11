package main

import "fmt"
//key与value进行交换
func main(){
	m1:=map[int]string{1:"a",2:"b",3:"c",4:"d"}
	fmt.Println(m1)
	m2:=make(map[string]int)
	for k,v:=range m1 {
		m2[v]=k
	}
	fmt.Println(m2)
}

