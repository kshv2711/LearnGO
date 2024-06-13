package main
import "fmt"
func main(){
	var number int
	fmt.Println("enter the number  to write a table")
	fmt.Scan(&number)
	i:=1
	for{
		if(i>10){
			break
		}
		fmt.Println(number,"x",i,"=",number*i)
		i++
	}
}