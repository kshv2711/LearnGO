package main
import "fmt"
func main(){
	l:=23
	b:=34
	func() {
		var area int
		area=l*b
		fmt.Println(area)
	}()

    func(a int,b int) {
	fmt.Println(a*b)
	}(20,23)
} 