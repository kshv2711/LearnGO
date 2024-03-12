package main

import (
	"fmt"
	"strings"
)
func main(){
	//join
	s:=[]string{"keshav","sharma","kanpur"}
	mystr:=strings.Join(s,"_")
	fmt.Printf("%T\n",mystr)
	fmt.Printf(mystr)
}