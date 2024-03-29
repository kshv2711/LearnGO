package main
import"fmt"
func SumAndOrder(A int) func(int) int{
	return func(B int) int {
		return A+B
	}
	
}
func main(){
	sum:=SumAndOrder(12)
	
	fmt.Println(sum(8))
}
// func Sum(v1,v2 int)int{
// 	return v1+v2

// }
// func SumAndOrder(A int) func(int) int{
// 	return func(B int) int {
// 		return Sum(A,B)
// 	}
	
// }
// func main(){
// 	sum:=SumAndOrder(12)
	
// 	fmt.Println(sum(8))
// }