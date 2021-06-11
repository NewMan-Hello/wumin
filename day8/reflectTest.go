package main
 import(
 	"fmt"
 	"reflect"
 )

func reflectTest(a interface{})  {
	//通过反射获取传入变量的typq、kind、值
	//获取 reflect.Type
	rTyp:=reflect.TypeOf(a)
	fmt.Println(rTyp)

	//获取reflect.Value
	rval:=reflect.ValueOf(a)
	fmt.Println(rval)

	//rval是reflec.Value类型
	n2:=2+rval.Int()
	fmt.Println("n2=",n2)
	fmt.Printf("rval=%v rval type=%T\n",rval,rval)

	//将rval转化为interface{}
	in:=rval.Interface()
	//将interface{}转化为变量，通过断言
	num2:=in.(int)
	fmt.Println("num2=",num2)

}

//结构体反射
func reflectTest2(b interface{}){
	//获取 reflect.Type
	rTyp:=reflect.TypeOf(b)
	fmt.Println(rTyp)

	//获取reflect.Value
	rval:=reflect.ValueOf(b)
	fmt.Println(rval)

	//将rval转化为interface{}
	in:=rval.Interface()
	fmt.Printf("in=%v in type=%T\n",in,in)
	//将interface{}转化为变量，通过断言
	stu,ok:=in.(student)
	if ok {
		fmt.Printf("stu=%v\n",stu.name)
	}

}

type student struct {
	name  string
	age int
}
func main(){
  //实现变量、interface{}、reflect.Value的相互转化
  var num int=100
  reflectTest(num)
  fmt.Println("==========================")

  stu:=student{
	  name: "123",
	  age:  18,
  }
  reflectTest2(stu)

}