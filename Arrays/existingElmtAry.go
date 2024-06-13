package main
import "fmt"
func main(){
	arr:=[]int{11,22,44,33,55,66,77,88,99,00}
	found:=false
	target:=33

	Index:=-1
	
	for i,value:=range arr{
		if(value==target){
			found=true
			Index=i
			break
		}
		

	}
	if found {
		fmt.Printf("ELEMENT ID Found ON [%d]=%d",Index,target)
		
	}else{fmt.Println("Not exist")}
}