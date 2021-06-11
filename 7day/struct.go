package main
import "fmt"
type humen struct {
	sex int
}
type person struct {
	humen
	name string
	age int
	//匿名结构，此content为一个字段，非结构体名称
	content struct{
		phone,city string
	}
}
 func main() {
   a:=&person{
   	name:"123",
   	age:18,
   	//嵌套赋值
   	humen:humen{sex:1},
   }
   //匿名结构体赋值
   a.content.phone="1349017077"
   //嵌套取值
   a.sex=0
   //a.name="123"
   //a.age=18
   fmt.Println(a)
   A(a)
   B(a)
   fmt.Println(a)
}
func A(per *person){
	per.age=13
	fmt.Println("A",per)
}
func B(per *person){
	per.age=15
	fmt.Println("B",per)
}