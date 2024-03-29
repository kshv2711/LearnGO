//A Higher-Order function is a function that receives a function as an argument or returns the function as output.
//Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.
package main
import"fmt"
func Increment(x int)func()int{
	return func() int {
		x++
		return x
	}

}
func main(){
	val:=Increment(9)
	val()
	fmt.Println(val())
}