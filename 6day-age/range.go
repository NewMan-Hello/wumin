package main

import (
	"fmt"
	"sort"
	)
func main(){
	//迭代
	sm:=make([]map[int]string,5)
	for i:=range sm{
		sm[i]=make(map[int]string,1)
		sm[i][1]="ok"
		fmt.Println(sm[i])
		i++
	}

	//排序
	m:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	s:=make([]int,len(m))
	v:=make([]string,len(m))
	i:=0
	for k,_:=range m{
		s[i]=k
		v[i]=m[s[i]]
		i++
	}
	fmt.Println(v)
	sort.Ints(s)
	fmt.Println(s)
}
