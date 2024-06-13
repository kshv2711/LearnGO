package main
import"fmt"
func Delete(slice []int, v int) []int{
	result:=[]int{}
	for _, value:=range slice{
		if value!=v {
			result=append(result, value)
		}
	}
	return result
}
func main(){
	Original:=[]int{99,88,77,66,55,44,33,22,66,11}
	target:=66
	Update:=Delete(Original,target)
	fmt.Println("original slice:",Original)
	fmt.Println("Updated slice:",Update)

}