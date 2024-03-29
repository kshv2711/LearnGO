package main
import "fmt"
func CallByReffence(a *int,b *int){
	var swap int
	swap=*a
	*a=*b
	*b=swap
	fmt.Println("after the value of a and b",*a,*b)
}
func main(){
	var c=34
	var d=43
	
	CallByReffence(&c,&d)
	fmt.Println("after the value of a and b",c,d)

}
