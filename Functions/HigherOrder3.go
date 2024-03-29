
package main
import "fmt"
func Operations(X int) func(int) func(int) int{
	return func(Y int) func(int) int {
		return func(Z int) int {
			return X*X+Y*Y+Z*Z
		}
	}
	
}
func main(){
	val:=Operations(23)
	
	fmt.Println(val(21)(31))
}