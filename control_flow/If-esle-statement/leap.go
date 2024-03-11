package main
import "fmt"
func main(){
	var year int
	var leap bool
	
	fmt.Println("Enter the Year ")
	fmt.Scanf("%d",&year)
	if year%400==0{
		leap=true
	}else if year%100==0{
		leap=false
	} else if year%4==0{
		leap=true}
	if leap {
		fmt.Println("this is leap")

	}else{
		fmt.Println("year is not leap")
	}
}