package main

import "fmt"
func main(){
	//slice
	//a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	//sa:=a[2:5]
	//fmt.Println(string(sa))
	//fmt.Println(len(sa),cap(sa))
	//sb:=sa[5:9]
	//fmt.Println(string(sb))

	//append
	//a:=[]int{1,2,3,4,5}
	//s1:=a[2:5]
	//s2:=a[1:3]
	//fmt.Println(s1,s2)
	////超出容量，重新分配数组
	//s2=append(s2,1,2,1,1,1,1,1,1,1)
	//s1[0]=9
	//fmt.Println(s1,s2)

	//copy
	s1:=[]int{1,2,3,4,5,6}
	s2:=[]int{7,8,9}
	copy(s1,s2)
	fmt.Println(s1)
	copy(s2,s1[3:6])
	fmt.Println(s2)
}
