package main
import "fmt"
type detail struct{
	name string
	phone int
	age  int
	salary float64

}
func updatedetails(d detail){
	d.name="keshav"
	d.phone= 9879899
	d.age=24
	d.salary=34234.232
	fmt.Println("at function")
	fmt.Println(d.age,d.name,d.phone,d.salary)

}
func updatedetailsUsingPointer(d *detail){
	d.name="keshav"
	d.phone= 9879899
	d.age=24
	d.salary=34234.232
	fmt.Println("at function")
	fmt.Println(d.age,d.name,d.phone,d.salary)

}


func main(){
	persent:=detail{
		name:"cheeku",
		phone:7437637,
		age:34,
		salary:43.434,
	}
	
	fmt.Println("before calling",persent)
	updatedetails(persent)
	// updatedetailsUsingPointer(&persent)
	fmt.Println("after calling",persent)


}