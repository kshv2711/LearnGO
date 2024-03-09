package main
import "fmt"
func poo() *int{
	x:=1
	return &x
}
func main(){
	var y *int
	y=poo()
	fmt.Printf("value of y is %d:",*y)



}