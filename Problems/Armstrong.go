package main
import "fmt"
func main(){
	var last,number,temp int
	var revert int=0
	fmt.Println("enter the number upto armstrong number")
	fmt.Scan(&number)
	temp=number
	for {
		last=temp%10
		revert+=last*last*last
		temp=temp/10

		if(temp==0){
			break
		}

	}
	if(revert==number){
		fmt.Print("number is armstrong=",revert)
    }else{
		fmt.Print("number is real number=",revert)
	}
}