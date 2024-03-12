package main

import (
	"fmt"
	"unicode/utf8"
)
func main(){
	mystr:="how can i iterate string"
	// using for loop
	// for i:=0;i<len(mystr);i++{
	// 	fmt.Printf("the index of i is [%d] and the value is [%c]\n",i,mystr[i])

	// }
	
	//using range loop
	// for i,v:=range mystr{
	// 	fmt.Printf("the index of i is [%d] and the value is [%c]\n",i,v)
	// }/
	println(len(mystr))
	n:=utf8.RuneCountInString(mystr)
	fmt.Println(n)

}