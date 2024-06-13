package main
import "fmt"
func main(){
	var orignal int=8
	oril:=&orignal
	a:=&oril
	b:=&a
	fmt.Printf("Address of original is %p\n",&orignal)
	fmt.Printf("Address of oril is %p\n",&oril)
	fmt.Printf("Address of oril is %v\n",oril)
	fmt.Println("the value of a is %v:",a)
	fmt.Println("the value of b is %v:",*b)
}