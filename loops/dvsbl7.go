package main
import "fmt"
func main(){
	num:=50
	for i := 1; i <=num; i++ {
		if i%7 !=0{
			
			continue
		}
		fmt.Printf("%d\n",i)
	}
}