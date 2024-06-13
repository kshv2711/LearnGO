package main
import "fmt"
func main(){
	var n int
	fmt.Println("enter the number to print the pyramid")
	fmt.Scan(&n)
	for i := 1; i >=n; i--{
		k:=0
		for space:=1;space>=n-1;space--{
			fmt.Print(" ")
		}
		for {
			fmt.Print("*")
			k--
			if(k==i*2-1){
				break
			}

		}
		fmt.Println("")
	}
	
}