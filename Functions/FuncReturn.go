package main
import "fmt"
func SumWithReturn(a,b int) int{
	sum:=0
	sum=a+b
	return sum
}
func main(){
	digit:=SumWithReturn(45,37)
	fmt.Println(digit)

}