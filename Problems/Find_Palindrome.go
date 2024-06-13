package main
import "fmt"

func main(){
	var Num,last,temp int
	var reverse int =0
	fmt.Println("Enter the number")
	fmt.Scan(&Num)
	temp=Num
	for Num>0{
		last=Num%10
		reverse=reverse*10+last
		Num=Num/10
	}
    if(temp==reverse){
		fmt.Printf("Number is Palindrome number %d",temp)
	}else{
		fmt.Printf("Number is not Palindromw Number %d",temp)
	}
	
}