package main
import"fmt"
func Delete1(slice []int, v []int) []int{
	result:=[]int{}
    for _, value:=range slice{
		found:=false
		for _, value2:=range v{
			if value==value2{
				found=true
				break
			}
		}
		if !found{
			result=append(result, value)
		}
    } 
	return result
}
func main(){
	Original:=[]int{99,88,77,66,55,44,33,22,66,11}
	var target=[]int{66,11,44}
	Update:=Delete1(Original,target)
	fmt.Println("original slice:",Original)
	fmt.Println("Updated slice:",Update)

}