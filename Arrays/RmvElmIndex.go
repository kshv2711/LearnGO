package main
import"fmt"
func Delete2(slice []int, k int) []int{
	result:=[]int{}
	
	for i, value:=range slice{
		if i!=k {
			result=append(result, value)
		}
		
	}
	return result
}
func main(){
	Original:=[]int{99,88,77,66,55,44,33,22,66,11}
	index:=2
	Update:=Delete2(Original,index)
	fmt.Println("original slice:",Original)
	fmt.Println("Updated slice:",Update)

}