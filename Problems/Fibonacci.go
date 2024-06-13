package main
import "fmt"
func main(){
	var n int
	num1:=0
	num2:=1
	fmt.Println("enter the number n")
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		if(i==1){
			fmt.Print(" ",num1)
			continue
		}
		if(i==2){
			fmt.Print(" ",num2)
		}
		newnum:=num1+num2
		num1=num2
		num2=newnum
		fmt.Print(" ",newnum)
	}
}