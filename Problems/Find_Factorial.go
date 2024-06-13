package main
import "fmt"
var fact uint64=1
var i int=1
var n int
func factorial(n int) uint64 {
	if(n<=0){
		fmt.Print("negative number or ZERO doesn't exist")
	}else{
		for i:=1;i<=n;i++{
			fact*=uint64(i)
		}
	}
	return fact
}
func main(){
	var num int
	fmt.Println("Enter the value to print factorial number")
	fmt.Scan(&num)
	fmt.Println(factorial(num))
}
