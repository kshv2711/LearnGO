package main

import (
	"fmt"
	"strings"
)
func PalindroneString(data string) bool{
	temp:=strings.ToLower(data)
	left, right:=0, len(temp)-1
	for left<right{
		if (temp[left]!=temp[right]) {
			return false
		}
		left++
		right--

	}
	return true

	 

}
func main(){
	
	fmt.Println(PalindroneString("mom"))
	
	
}