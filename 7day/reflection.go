package main

import ("fmt"
"reflect"
)
type User struct {
	Id  int
	Name  string
	Age int

}
func (u User) Hello(){
	fmt.Print("hello world")
}
func main(){
	u:=User{1,"123",18}
	info(u)

}
func info(o interface{}){
	t:=reflect.TypeOf(o)
	fmt.Println("type:",t.Name())

	v:=reflect.ValueOf(o)
	fmt.Println("fields:")
	for i:=0;i<t.NumField() ;i++  {
		f:=t.Field(i)
		val:=v.Field(i).Interface()
		fmt.Println(f.Name,f.Type,val)

	}
	for i:=0;i<t.NumMethod() ;i++  {
		m:=t.Method(i)
		fmt.Println(m.Name,m.Type)
	}
}
