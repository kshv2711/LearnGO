package main

import (
	"fmt"
	"strings"
)
func main(){
	str1:="keshav"
	f:=strings.Fields(str1)
	for i,j:=0, len(f)-1; i<j; i, j =i+1,j-1 {
		f[i], f[j]= f[j], f[i]
	}
	fmt.Println(strings.Join(f," "))

}