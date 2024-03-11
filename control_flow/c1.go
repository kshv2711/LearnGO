package main
import "fmt"
func main(){
	var i int
	fmt.Println("enter your age please")
	fmt.Scan(&i)
	if i<18 {
		rg := 18-i
		fmt.Println("you cannot cast the vote at age of",i ,"you need more at atleast", rg, "years")
	} else {
		fmt.Printf("You can the cast the vote at  age of %d",i)
	}
}