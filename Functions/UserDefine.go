package main
import"fmt"
 type first func(int) int
 type second func(int) first
func anotherOperations(x int) second{
	return func( y int) first {
		return func(z int) int {
			return x*x+y*y+z*z

		}
	}
}
func main(){
	value:=anotherOperations(12)
	fmt.Println(value(12)(13))

}