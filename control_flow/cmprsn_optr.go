package main
import "fmt"
func main(){
	var speed int =100
	fast:= speed>60
	slow:= speed<40
	top:= speed==90

	fmt.Printf("the seep of vehicle :%t\t",fast)
	fmt.Printf("\nthe seep of vehicle :%t\t",slow)
	fmt.Printf("\nthe seep of vehicle :%t\t",top)

} 