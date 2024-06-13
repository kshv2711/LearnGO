package main
import "fmt"
func main(){
	Arr2:=[10]int{1,2,3,4,5,6,7,8,9,0}
	Arr1:=Arr2
	left:=0
	right:=len(Arr1)-1
	for left<right{
		Arr1[left],Arr1[right]=Arr1[right],Arr1[left]
		left++
		right--
	}
	fmt.Println(Arr1)
}