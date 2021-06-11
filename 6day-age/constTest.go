package main
import(
	"fmt"
)
const a int=1
const b,c,d=1,"2",3
const(
	e=1
	f="123"
	g=len(f)
)
func main(){
	fmt.Println(e,f,g)
}
