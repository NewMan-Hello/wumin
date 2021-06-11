package main

import "fmt"

type USB interface {
	Name( ) string
	//嵌套接口
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}
func (pc PhoneConnecter) Name() string{
	return pc.name
}
func (pc PhoneConnecter) Connect()  {
	fmt.Print("connecter:",pc.name)
}
func main(){
	a:=PhoneConnecter{name:"12300"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb interface{}){
	switch v:=usb.(type) {
	case PhoneConnecter:
		fmt.Print("Disconnect:",v.name)
	default:
		fmt.Print("UnKown decive")
	}
}