package main

import(
	"fmt"
)
func main(){
	s1:=[]int{1,2,3,4}
	a:=1
	b:=D
	A(s1)
	fmt.Println(s1)
	B(a)
	fmt.Println(a)
	C(&a)
	fmt.Println(a)
	b()

	f:=closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))


}
//引用类型修改的是地址
func A(s []int){
	s[0]=5
	s[1]=6
	s[2]=7
	s[3]=8
	fmt.Println(s)
}
//值类型修改的是值
func B(a int)  {
	a=2
	fmt.Println(a)
}
//通过指针修改地址
func C(a *int)  {
	*a=2
	fmt.Println(*a)
}
//函数本身也是类型
func D(){
	fmt.Println("Func D")
}

//闭包
func closure(x int) func(int)int{
	fmt.Println(&x)
	return func(y int)int{
		fmt.Println(&x)
		return x+y
	}

}