package main

import "fmt"
func main(){
	//a:=[...]int{19:1}
	//fmt.Println(a)
	//b:=[10]int{}
	//b[1]=1
	//fmt.Println(b)
	//c:=new([10]int)
	//c[1]=1
	//fmt.Println(c)

	//冒泡排序
	a:=[...]int {5,3,6,3,9}
	fmt.Println(a)
	num:=len(a)
	for i:=0;i<num;i++{
		for j:=i+1;j<num ;j++  {
			if a[i]<a[j]{
				temp:=a[i]
				a[i]=a[j]
				a[j]=temp
			}

		}
	}
	fmt.Println(a)

}