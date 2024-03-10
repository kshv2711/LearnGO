package main
import "fmt"
func main(){
	arry1:=[]int{11,22,33,44,55,66,77,88,99,00,}
	arry2:=&arry1

	fmt.Println("original array arry1",arry1)
	fmt.Println("copied array arry2",*arry2)
}