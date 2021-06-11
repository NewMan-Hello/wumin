package main
import "fmt"
type A struct {
	name string
}
type B struct {
	name string
}
func main(){
	a:=&A{}
	a.Print()
	fmt.Println(a.name)
	b:=B{}
	b.Print()
	fmt.Println(b.name)

}
func (a *A) Print(){
	a.name="AA"
	fmt.Println("A")
}
func (b B) Print()  {
	b.name="BB"
	fmt.Println("B")
}