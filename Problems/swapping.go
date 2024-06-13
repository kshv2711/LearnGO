package main
import"fmt"
func main(){
	var Dig1, Dig2 int
	fmt.Println("enter the nuber for swapping")
	fmt.Scanln(&Dig1,&Dig2)
	fmt.Println("Before",Dig1,Dig2)
	Dig1=Dig1+Dig2
	Dig2=Dig1-Dig2
	Dig1=Dig1-Dig2
	fmt.Println("aftr the swapping",Dig1,Dig2)
}