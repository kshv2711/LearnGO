package main
import "fmt"
func main(){
	var number  int
	sum:=0
	fmt.Println("enter the number for Upto sum")
	fmt.Scanln(&number)
	for i:=0;i<=number;i++{
		sum+=i
	}
	fmt.Println(sum)
}