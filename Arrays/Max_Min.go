package main
import "fmt"
func main(){
	arr:=[10]int{1,8,9,5,4,13,0}
	max:=arr[0]
	min:=arr[0]
	for _,value:=range arr{
		if value>max{
			max=value
		}
		if value<min{
			min=value
		}
	}
	fmt.Printf("the max value is=%d\n",max)
	fmt.Printf("the mix value is=%d",min)


}