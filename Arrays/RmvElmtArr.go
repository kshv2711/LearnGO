package main

import (
	"fmt"
	
)
func main(){
	Arr:=[]int{22,33,44,44,44,77,77,88}
	slices:=append(Arr[ :2], Arr[3:]...)
	fmt.Print(slices)
}