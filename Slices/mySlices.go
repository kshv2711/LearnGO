package main

import (
	"fmt"
	
)
func main(){
	number:=[]int{1,2,3,4,5,6,7,8,9}

	//using  for loop
	for i:=0;i<len(number);i++{
		fmt.Println(number[i])
	}
}