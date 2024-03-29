package main 
import "fmt"
func EmpolyeeDetail(name string,age int,salary float64)(string,float64,int){

	var avgSalAge float64
	avgSalAge=(salary/float64(age))
	return name, avgSalAge,age// Return statement with specify variable name
}
func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}


func main(){
	name,avgSal,age:=EmpolyeeDetail("keshav",24,20000000.)
	fmt.Println(name,avgSal,age)
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}

