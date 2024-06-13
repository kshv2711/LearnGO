package main

import (
	"fmt"
	"sort"
)
func main(){
	arr1:=[]int{88,66,55,44,6677,44,55,22,90}
	Slice:=arr1[:]
	sort.Ints(Slice)
	fmt.Println(Slice)
}