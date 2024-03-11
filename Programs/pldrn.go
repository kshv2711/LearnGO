package main
import "fmt"
func main(){
	number:=123321
	num:=number
	var revers  int
	for num>0 {
		lDigt:=num%10
		revers=revers*10+lDigt
		num=num/10
	} 
	if number == revers{
		fmt.Println("number is palindrone")
	} else{
		fmt.Println("number is not a palindrone")
	}
}