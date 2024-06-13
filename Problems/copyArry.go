package main
import "fmt"
func main(){
	arr1:=[]int{1,2,3,4,5,6,8,9,0}
	arr2:=make([]int,len(arr1))
	n:=copy(arr2,arr1)
	fmt.Println(arr2,arr1,n)

}