package main
import "fmt"
func main(){
	value:=100;
	dValue:=&value
	fmt.Println(&dValue)//address of value
	fmt.Println(*dValue)//value of address of dvalue 
}