package main
import "fmt"
func SumAndProduct(a ...float64)(float64,float64){
	sum:=0.
	product:=1.
	for _,v := range a{
		sum +=v 
		product *=v
	}
	return sum,product

} 
func main(){
	sum, prod:=SumAndProduct(3,3)
	fmt.Println(sum,prod)
}
