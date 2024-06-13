package main
import "fmt"
func main(){
	var Num[100] int
	var n int
	fmt.Println("enter the number of  element in array")
	fmt.Scanln(&n)
	for i:=0;i<n;i++{
		fmt.Println("the the element in array")
		fmt.Scanln(&Num[i])
	}
	for j:=1;j<n;j++{
		if(Num[0]<Num[j]){
			Num[0]=Num[j]
		}
	}
	print(Num[0])

}